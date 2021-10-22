//go:build !simple
// +build !simple

package command

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mzlog "github.com/asim/go-micro/plugins/logger/zerolog/v3"
	"github.com/rs/zerolog"
	"go-micro.dev/v4/logger"

	"gopkg.in/yaml.v2"

	"github.com/owncloud/ocis/ocis-pkg/config"
	pkglog "github.com/owncloud/ocis/ocis-pkg/log"
	"github.com/owncloud/ocis/ocis/pkg/register"
	pcommand "github.com/owncloud/ocis/proxy/pkg/command"
	proxyFlagset "github.com/owncloud/ocis/proxy/pkg/flagset"
	"github.com/thejerf/suture/v4"
	"github.com/urfave/cli/v2"
)

type serviceFuncMap map[string]func(*config.Config) suture.Service

type Service struct {
	Supervisor *suture.Supervisor
	Services   serviceFuncMap
	Log        log.Logger

	serviceToken map[string][]suture.ServiceToken
	context      context.Context
	cancel       context.CancelFunc
}

var (
	// tolerance controls backoff cycles from the supervisor
	tolerance = 5

	// totalBackoff stops the retries after the tolerance has been reached
	totalBackoff = 0
)

func parseConfig(cfg *config.Config) error {
	// only deal with yaml to make matters easier. Will change
	b, _ := os.ReadFile("/Users/aunger/.ocis/runtime/ocis.yaml")
	err := yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return nil
}

// Experimental uses an opinionated experimental runtime.
func Experimental(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "experimental",
		Usage:    "Start fullstack server",
		Category: "Experimental",
		Flags:    proxyFlagset.ServerWithConfig(cfg.Proxy), // TODO(refs) must not use flagsets during supervised mode
		Before: func(c *cli.Context) error {
			// before parsing config use a default logger to log any events
			l := pkglog.NewLogger().Level(1) // info

			if _, err := os.Stat("/Users/aunger/.ocis/runtime/ocis.yaml"); err != nil {
				l.Info().Msg("config file not set in default location, continuing without config...")
				return nil
			}

			// after config parsing, there should be logging info in the config, so use it to initialize loggers
			if err := parseConfig(cfg); err != nil {
				return err
			}

			return nil
		},
		Action: func(c *cli.Context) error {
			var (
				s               = Service{}
				rootCtx, cancel = context.WithCancel(context.Background())
			)
			defer cancel()

			// halt listens for interrupt signals and blocks.
			halt := make(chan os.Signal, 1)
			signal.Notify(halt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

			// prepare supervisor
			s.Supervisor = suture.New("ocis", suture.Spec{
				EventHook: func(e suture.Event) {
					if e.Type() == suture.EventTypeBackoff {
						totalBackoff++
						if totalBackoff == tolerance {
							halt <- os.Interrupt
						}
					}
				},
				FailureThreshold: 5,
				FailureBackoff:   3 * time.Second,
			})

			// setup runtime net listener
			l, err := net.Listen("tcp", net.JoinHostPort("localhost", "16666"))
			if err != nil {
				panic(err)
			}

			// prevents from undesired logging from go-micro
			setMicroLogger()

			s.Supervisor.Add(pcommand.NewPSuture(rootCtx, *cfg))

			// run supervised services
			go s.Supervisor.ServeBackground(s.context)

			go func() {
				<-halt
				os.Exit(0)
			}()

			return http.Serve(l, nil)
		},
	}
}

func init() {
	register.AddCommand(Experimental)
}

// for logging reasons we don't want the same logging level on both oCIS and micro. As a framework builder we do not
// want to expose to the end user the internal framework logs unless explicitly specified.
func setMicroLogger() {
	if os.Getenv("MICRO_LOG_LEVEL") == "" {
		_ = os.Setenv("MICRO_LOG_LEVEL", "error")
	}

	lev, err := zerolog.ParseLevel(os.Getenv("MICRO_LOG_LEVEL"))
	if err != nil {
		lev = zerolog.ErrorLevel
	}
	logger.DefaultLogger = mzlog.NewLogger(logger.WithLevel(logger.Level(lev)))
}
