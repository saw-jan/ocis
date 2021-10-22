package command

import (
	"context"
	"fmt"

	"github.com/thejerf/suture/v4"

	ociscfg "github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/proxy/pkg/metrics"
	"github.com/owncloud/ocis/proxy/pkg/proxy"
	proxyHTTP "github.com/owncloud/ocis/proxy/pkg/server/http"

	"github.com/oklog/run"
	"github.com/owncloud/ocis/proxy/pkg/config"
	"github.com/urfave/cli/v2"
)

// Action does no config parsing. It assumes the config it gets is definitive.
func Action(cfg config.Config) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		var (
			g           = run.Group{}
			ctx, cancel = context.WithCancel(context.Background())
			logger      = NewLogger(&cfg)
		)

		defer cancel()

		rp := proxy.NewMultiHostReverseProxy(proxy.Logger(logger), proxy.Config(&cfg))

		server, err := proxyHTTP.Server(
			proxyHTTP.Handler(rp),
			proxyHTTP.Logger(logger),
			proxyHTTP.Context(ctx),
			proxyHTTP.Config(&cfg),
			proxyHTTP.Metrics(metrics.New()),
			proxyHTTP.Middlewares(loadMiddlewares(ctx, logger, &cfg)),
		)

		if err != nil {
			logger.Error().Err(err).Str("server", "http").Msg("Failed to initialize server")
			return err
		}

		g.Add(func() error {
			return server.Run()
		}, func(_ error) {
			logger.Info().Str("server", "http").Msg("Shutting down server")
			cancel() // calling cancel() will cause the micro-service to shut down.
		})

		return g.Run()
	}
}

type PSuture struct {
	ctx context.Context
	cfg config.Config
}

// NewPSuture creates a new suture.Service.
// The only real shared config should be about logging
func NewPSuture(ctx context.Context, cfg ociscfg.Config) suture.Service {
	c := configureLogging(cfg)

	return PSuture{
		ctx: ctx,
		cfg: c,
	}
}

func (s PSuture) Serve(ctx context.Context) error {
	s.cfg.Context = ctx
	if err := Action(s.cfg)(nil); err != nil {
		return fmt.Errorf("you're in cahoots")
	}

	return nil
}

func configureLogging(cfg ociscfg.Config) config.Config {
	cfg.Proxy.Log.Level = cfg.Log.Level
	cfg.Proxy.Log.Color = cfg.Log.Color
	cfg.Proxy.Log.Pretty = cfg.Log.Pretty

	return *cfg.Proxy
}
