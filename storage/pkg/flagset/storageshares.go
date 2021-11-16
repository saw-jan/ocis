package flagset

import (
	"github.com/owncloud/ocis/ocis-pkg/flags"
	"github.com/owncloud/ocis/storage/pkg/config"
	"github.com/urfave/cli/v2"
)

// StorageSharesWithConfig applies cfg to the root flagset
func StorageSharesWithConfig(cfg *config.Config) []cli.Flag {
	flags := []cli.Flag{

		// debug ports are the odd ports
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       flags.OverrideDefaultString(cfg.Reva.StorageShares.DebugAddr, "127.0.0.1:9156"),
			Usage:       "Address to bind debug server",
			EnvVars:     []string{"STORAGE_SHARES_DEBUG_ADDR"},
			Destination: &cfg.Reva.StorageShares.DebugAddr,
		},

		// Services

		// Share Storage Provider

		&cli.StringFlag{
			Name:        "grpc-network",
			Value:       flags.OverrideDefaultString(cfg.Reva.StorageShares.GRPCNetwork, "tcp"),
			Usage:       "Network to use for the storage service, can be 'tcp', 'udp' or 'unix'",
			EnvVars:     []string{"STORAGE_SHARES_GRPC_NETWORK"},
			Destination: &cfg.Reva.StorageShares.GRPCNetwork,
		},
		&cli.StringFlag{
			Name:        "grpc-addr",
			Value:       flags.OverrideDefaultString(cfg.Reva.StorageShares.GRPCAddr, "127.0.0.1:9154"),
			Usage:       "Address to bind storage service",
			EnvVars:     []string{"STORAGE_SHARES_GRPC_ADDR"},
			Destination: &cfg.Reva.StorageShares.GRPCAddr,
		},
		&cli.BoolFlag{
			Name:        "read-only",
			Value:       flags.OverrideDefaultBool(cfg.Reva.StorageUsers.ReadOnly, false),
			Usage:       "use storage driver in read-only mode",
			EnvVars:     []string{"STORAGE_USERS_READ_ONLY", "OCIS_STORAGE_READ_ONLY"},
			Destination: &cfg.Reva.StorageShares.ReadOnly,
		},

		// FIXME currently the sharesstorageprovider directly talks to the user share provider

		// Gateway

		&cli.StringFlag{
			Name:        "reva-gateway-addr",
			Value:       flags.OverrideDefaultString(cfg.Reva.Gateway.Endpoint, "127.0.0.1:9142"),
			Usage:       "Address of REVA gateway endpoint",
			EnvVars:     []string{"REVA_GATEWAY"},
			Destination: &cfg.Reva.Gateway.Endpoint,
		},

		// User share provider

		&cli.StringFlag{
			Name:        "sharing-endpoint",
			Value:       flags.OverrideDefaultString(cfg.Reva.Sharing.Endpoint, "localhost:9150"),
			Usage:       "endpoint to use for the storage service",
			EnvVars:     []string{"STORAGE_SHARING_ENDPOINT"},
			Destination: &cfg.Reva.Sharing.Endpoint,
		},
	}

	flags = append(flags, TracingWithConfig(cfg)...)
	flags = append(flags, DebugWithConfig(cfg)...)
	flags = append(flags, SecretWithConfig(cfg)...)

	return flags
}
