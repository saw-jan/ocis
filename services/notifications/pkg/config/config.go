// Package config provides the service configuration.
package config

import (
	"context"

	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
)

// Config combines all available configuration parts.
type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing"`
	Log     *Log     `yaml:"log"`
	Debug   Debug    `yaml:"debug"`

	WebUIURL string `yaml:"ocis_url" env:"OCIS_URL;NOTIFICATIONS_WEB_UI_URL" desc:"The public facing URL of the oCIS Web UI, used e.g. when sending notification eMails"`

	Notifications  Notifications        `yaml:"notifications"`
	GRPCClientTLS  shared.GRPCClientTLS `yaml:"grpc_client_tls"`
	ServiceAccount ServiceAccount       `yaml:"service_account"`

	Context context.Context `yaml:"-"`
}

// Notifications defines the config options for the notifications service.
type Notifications struct {
	SMTP              SMTP                  `yaml:"SMTP"`
	Events            Events                `yaml:"events"`
	EmailTemplatePath string                `yaml:"email_template_path" env:"OCIS_EMAIL_TEMPLATE_PATH;NOTIFICATIONS_EMAIL_TEMPLATE_PATH" desc:"Path to Email notification templates overriding embedded ones."`
	TranslationPath   string                `yaml:"translation_path" env:"OCIS_TRANSLATION_PATH;NOTIFICATIONS_TRANSLATION_PATH" desc:"(optional) Set this to a path with custom translations to overwrite the builtin translations. Note that file and folder naming rules apply, see the documentation for more details."`
	DefaultLanguage   string                `yaml:"default_language" env:"OCIS_DEFAULT_LANGUAGE" desc:"The default language used by services and the WebUI. If not defined, English will be used as default. See the documentation for more details."`
	RevaGateway       string                `yaml:"reva_gateway" env:"OCIS_REVA_GATEWAY" desc:"CS3 gateway used to look up user metadata"`
	GRPCClientTLS     *shared.GRPCClientTLS `yaml:"grpc_client_tls"`
}

// SMTP combines the smtp configuration options.
type SMTP struct {
	Host           string `yaml:"smtp_host" env:"NOTIFICATIONS_SMTP_HOST" desc:"SMTP host to connect to."`
	Port           int    `yaml:"smtp_port" env:"NOTIFICATIONS_SMTP_PORT" desc:"Port of the SMTP host to connect to."`
	Sender         string `yaml:"smtp_sender" env:"NOTIFICATIONS_SMTP_SENDER" desc:"Sender address of emails that will be sent (e.g. 'ownCloud <noreply@example.com>'."`
	Username       string `yaml:"smtp_username" env:"NOTIFICATIONS_SMTP_USERNAME" desc:"Username for the SMTP host to connect to."`
	Password       string `yaml:"smtp_password" env:"NOTIFICATIONS_SMTP_PASSWORD" desc:"Password for the SMTP host to connect to."`
	Insecure       bool   `yaml:"insecure" env:"NOTIFICATIONS_SMTP_INSECURE" desc:"Allow insecure connections to the SMTP server."`
	Authentication string `yaml:"smtp_authentication" env:"NOTIFICATIONS_SMTP_AUTHENTICATION" desc:"Authentication method for the SMTP communication. Possible values are 'login', 'plain', 'crammd5', 'none' or 'auto'. If set to 'auto' or unset, the authentication method is automatically negotiated with the server."`
	Encryption     string `yaml:"smtp_encryption" env:"NOTIFICATIONS_SMTP_ENCRYPTION" desc:"Encryption method for the SMTP communication. Possible values  are 'starttls', 'ssl', 'ssltls', 'tls'  and 'none'." deprecationVersion:"5.0.0" removalVersion:"6.0.0" deprecationInfo:"The NOTIFICATIONS_SMTP_ENCRYPTION values 'ssl' and 'tls' are deprecated and will be removed in the future." deprecationReplacement:"Use 'starttls' instead of 'tls' and 'ssltls' instead of 'ssl'."`
}

// Events combines the configuration options for the event bus.
type Events struct {
	Endpoint             string `yaml:"endpoint" env:"OCIS_EVENTS_ENDPOINT;NOTIFICATIONS_EVENTS_ENDPOINT" desc:"The address of the event system. The event system is the message queuing service. It is used as message broker for the microservice architecture."`
	Cluster              string `yaml:"cluster" env:"OCIS_EVENTS_CLUSTER;NOTIFICATIONS_EVENTS_CLUSTER" desc:"The clusterID of the event system. The event system is the message queuing service. It is used as message broker for the microservice architecture. Mandatory when using NATS as event system."`
	TLSInsecure          bool   `yaml:"tls_insecure" env:"OCIS_INSECURE;NOTIFICATIONS_EVENTS_TLS_INSECURE" desc:"Whether to verify the server TLS certificates."`
	TLSRootCACertificate string `yaml:"tls_root_ca_certificate" env:"OCIS_EVENTS_TLS_ROOT_CA_CERTIFICATE;NOTIFICATIONS_EVENTS_TLS_ROOT_CA_CERTIFICATE" desc:"The root CA certificate used to validate the server's TLS certificate. If provided NOTIFICATIONS_EVENTS_TLS_INSECURE will be seen as false."`
	EnableTLS            bool   `yaml:"enable_tls" env:"OCIS_EVENTS_ENABLE_TLS;NOTIFICATIONS_EVENTS_ENABLE_TLS" desc:"Enable TLS for the connection to the events broker. The events broker is the ocis service which receives and delivers events between the services.."`
	AuthUsername         string `yaml:"username" env:"OCIS_EVENTS_AUTH_USERNAME;NOTIFICATIONS_EVENTS_AUTH_USERNAME" desc:"The username to authenticate with the events broker. The events broker is the ocis service which receives and delivers events between the services.."`
	AuthPassword         string `yaml:"password" env:"OCIS_EVENTS_AUTH_PASSWORD;NOTIFICATIONS_EVENTS_AUTH_PASSWORD" desc:"The password to authenticate with the events broker. The events broker is the ocis service which receives and delivers events between the services.."`
}

// ServiceAccount is the configuration for the used service account
type ServiceAccount struct {
	ServiceAccountID     string `yaml:"service_account_id" env:"OCIS_SERVICE_ACCOUNT_ID;NOTIFICATIONS_SERVICE_ACCOUNT_ID" desc:"The ID of the service account the service should use. See the 'auth-service' service description for more details."`
	ServiceAccountSecret string `yaml:"service_account_secret" env:"OCIS_SERVICE_ACCOUNT_SECRET;NOTIFICATIONS_SERVICE_ACCOUNT_SECRET" desc:"The service account secret."`
}
