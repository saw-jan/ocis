package command

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/cs3org/reva/v2/cmd/revad/runtime"
	"github.com/gofrs/uuid"
	"github.com/oklog/run"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	"github.com/owncloud/ocis/v2/ocis-pkg/registry"
	"github.com/owncloud/ocis/v2/ocis-pkg/sync"
	"github.com/owncloud/ocis/v2/ocis-pkg/tracing"
	"github.com/owncloud/ocis/v2/ocis-pkg/version"
	"github.com/owncloud/ocis/v2/services/sharing/pkg/config"
	"github.com/owncloud/ocis/v2/services/sharing/pkg/config/parser"
	"github.com/owncloud/ocis/v2/services/sharing/pkg/logging"
	"github.com/owncloud/ocis/v2/services/sharing/pkg/revaconfig"
	"github.com/owncloud/ocis/v2/services/sharing/pkg/server/debug"
	"github.com/urfave/cli/v2"
)

// Server is the entry point for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start the %s service without runtime (unsupervised mode)", cfg.Service.Name),
		Category: "server",
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			logger := logging.Configure(cfg.Service.Name, cfg.Log)
			tracingProvider, err := tracing.GetServiceTraceProvider(cfg.Tracing, cfg.Service.Name)
			if err != nil {
				return err
			}
			gr := run.Group{}
			ctx, cancel := defineContext(cfg)

			defer cancel()

			// precreate folders
			if cfg.UserSharingDriver == "json" && cfg.UserSharingDrivers.JSON.File != "" {
				if err := os.MkdirAll(filepath.Dir(cfg.UserSharingDrivers.JSON.File), os.FileMode(0700)); err != nil {
					return err
				}
			}
			if cfg.PublicSharingDriver == "json" && cfg.PublicSharingDrivers.JSON.File != "" {
				if err := os.MkdirAll(filepath.Dir(cfg.PublicSharingDrivers.JSON.File), os.FileMode(0700)); err != nil {
					return err
				}
			}

			gr.Add(func() error {
				pidFile := path.Join(os.TempDir(), "revad-"+cfg.Service.Name+"-"+uuid.Must(uuid.NewV4()).String()+".pid")
				rCfg, err := revaconfig.SharingConfigFromStruct(cfg, logger)
				if err != nil {
					return err
				}
				reg := registry.GetRegistry()

				runtime.RunWithOptions(rCfg, pidFile,
					runtime.WithLogger(&logger.Logger),
					runtime.WithRegistry(reg),
					runtime.WithTraceProvider(tracingProvider),
				)

				return nil
			}, func(err error) {
				logger.Error().
					Err(err).
					Str("server", cfg.Service.Name).
					Msg("Shutting down server")

				cancel()
				os.Exit(1)
			})

			debugServer, err := debug.Server(
				debug.Logger(logger),
				debug.Context(ctx),
				debug.Config(cfg),
			)

			if err != nil {
				logger.Info().Err(err).Str("server", "debug").Msg("Failed to initialize server")
				return err
			}

			gr.Add(debugServer.ListenAndServe, func(_ error) {
				cancel()
			})

			if !cfg.Supervised {
				sync.Trap(&gr, cancel)
			}

			grpcSvc := registry.BuildGRPCService(cfg.GRPC.Namespace+"."+cfg.Service.Name, uuid.Must(uuid.NewV4()).String(), cfg.GRPC.Addr, version.GetString())
			if err := registry.RegisterService(ctx, grpcSvc, logger); err != nil {
				logger.Fatal().Err(err).Msg("failed to register the grpc service")
			}

			return gr.Run()
		},
	}
}

// defineContext sets the context for the service. If there is a context configured it will create a new child from it,
// if not, it will create a root context that can be cancelled.
func defineContext(cfg *config.Config) (context.Context, context.CancelFunc) {
	return func() (context.Context, context.CancelFunc) {
		if cfg.Context == nil {
			return context.WithCancel(context.Background())
		}
		return context.WithCancel(cfg.Context)
	}()
}
