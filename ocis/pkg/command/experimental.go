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

	"github.com/mohae/deepcopy"

	"github.com/owncloud/ocis/ocis-pkg/config"
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

// Experimental uses an opinionated experimental runtime.
func Experimental(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "experimental",
		Usage:    "Start fullstack server",
		Category: "Experimental",
		Flags:    proxyFlagset.ServerWithConfig(cfg.Proxy),
		Action: func(c *cli.Context) error {
			var (
				s               = Service{}
				rootCtx, cancel = context.WithCancel(context.Background())
			)
			defer cancel()

			// halt listens for interrupt signals and blocks.
			halt := make(chan os.Signal, 1)
			signal.Notify(halt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

			// tolerance controls backoff cycles from the supervisor.
			tolerance := 5
			totalBackoff := 0

			// Start creates its own supervisor. Running services under `ocis server` will create its own supervision tree.
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

			// add the proxy service
			s.Supervisor.Add(pcommand.NewPSuture(rootCtx, deepcopy.Copy(*cfg.Proxy)))

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
