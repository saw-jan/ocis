// Package config contains the configuration for the ocis-thumbnails service
package config

import (
	"context"

	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
	"go-micro.dev/v4/client"
)

// Config combines all available configuration parts.
type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing"`
	Log     *Log     `yaml:"log"`
	Debug   Debug    `yaml:"debug"`

	GRPC GRPCConfig `yaml:"grpc"`
	HTTP HTTP       `yaml:"http"`

	GRPCClientTLS *shared.GRPCClientTLS `yaml:"grpc_client_tls"`
	GrpcClient    client.Client         `yaml:"-"`

	Thumbnail Thumbnail `yaml:"thumbnail"`

	Context context.Context `yaml:"-"`
}

// FileSystemStorage defines the available filesystem storage configuration.
type FileSystemStorage struct {
	RootDirectory string `yaml:"root_directory" env:"THUMBNAILS_FILESYSTEMSTORAGE_ROOT" desc:"The directory where the filesystem storage will store the thumbnails. If not defined, the root directory derives from $OCIS_BASE_DATA_PATH:/thumbnails."`
}

// Thumbnail defines the available thumbnail related configuration.
type Thumbnail struct {
	Resolutions         []string          `yaml:"resolutions" env:"THUMBNAILS_RESOLUTIONS" desc:"The supported list of target resolutions in the format WidthxHeight like 32x32. You can define any resolution as required. See the Environment Variable Types description for more details."`
	FileSystemStorage   FileSystemStorage `yaml:"filesystem_storage"`
	WebdavAllowInsecure bool              `yaml:"webdav_allow_insecure" env:"OCIS_INSECURE;THUMBNAILS_WEBDAVSOURCE_INSECURE" desc:"Ignore untrusted SSL certificates when connecting to the webdav source."`
	CS3AllowInsecure    bool              `yaml:"cs3_allow_insecure" env:"OCIS_INSECURE;THUMBNAILS_CS3SOURCE_INSECURE" desc:"Ignore untrusted SSL certificates when connecting to the CS3 source."`
	RevaGateway         string            `yaml:"reva_gateway" env:"OCIS_REVA_GATEWAY" desc:"CS3 gateway used to look up user metadata"`
	FontMapFile         string            `yaml:"font_map_file" env:"THUMBNAILS_TXT_FONTMAP_FILE" desc:"The path to a font file for txt thumbnails."`
	TransferSecret      string            `yaml:"transfer_secret" env:"THUMBNAILS_TRANSFER_TOKEN" desc:"The secret to sign JWT to download the actual thumbnail file."`
	DataEndpoint        string            `yaml:"data_endpoint" env:"THUMBNAILS_DATA_ENDPOINT" desc:"The HTTP endpoint where the actual thumbnail file can be downloaded."`
}
