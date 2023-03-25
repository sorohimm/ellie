package config

import (
	"github.com/sorohimm/ellie/internal/storage/postgres"
	"time"
)

type Config struct {
	HTTP     *HTTPConfig      `group:"http option" namespace:"http" env-namespace:"HTTP"`
	Log      *Logger          `group:"logger option" namespace:"log" env-namespace:"LOG"`
	Source   *SourceConfig    `group:"source option" namespace:"source" env-namespace:"SOURCE"`
	API      *ApiConfig       `group:"api option" namespace:"api" env-namespace:"API"`
	CORS     *CORS            `group:"cors options" namespace:"cors" env-namespace:"CORS"`
	Postgres *postgres.Config `group:"pgx pool option" namespace:"postgres" env-namespace:"POSTGRES"`
}

type Logger struct {
	Level   string `short:"l" long:"level" env:"LEVEL" description:"logging level" default:"DEBUG"`
	EncType string `long:"enctype" env:"ENCTYPE" description:"log as json or not (console|json)" default:"json" `
}

type CORS struct {
	Enabled bool   `long:"enabled"  env:"ENABLED" description:"enable or disable CORS OPTIONS request processing (if false - all origins allowed otherwise only specified origins allowed )"`
	Origins string `long:"origins" env:"ORIGINS" default:"http://localhost,http://localhost:*,http://*.megafon.ru,http://*.megafon.ru:*" description:"comma separated list of allowed origins (default value the same as if all origins has been allowed)" `
}

func (o *CORS) SplitOrigins() []string {
	return splitStringBy(o.Origins, ",", " \t\r\n")
}

type SourceConfig struct {
	Client *ClientParams `group:"http client option" namespace:"client" env-namespace:"CLIENT"`
}

type ApiConfig struct {
	Root        string `long:"path-root" env:"PATH_ROOT" default:"eapi"`
	ServiceName string `long:"path-service-name" env:"PATH_SERVICE_NAME" default:"pythia-apigw"`
	ApiVersion  string `long:"api-version" env:"PATH_VERSION" default:"v1"`
}

type ClientParams struct {
	MaxIdleConns        int           `long:"maxIdleConns" default:"10"  env:"MAX_IDLE_CONNS" description:"the maximum number of idle (keep-alive) connections"`
	MaxConnsPerHost     int           `long:"MaxConnsPerHost" default:"10"  env:"MAX_CONNS_PER_HOST" description:"optionally limits the total number of connections per host"`
	MaxIdleConnsPerHost int           `long:"MaxIdleConnsPerHost" default:"10"  env:"MAX_IDLE_CONNS_PER_HOST" description:"if non-zero, controls the maximum idle (keep-alive) connections to keep per-host"`
	Timeout             time.Duration `long:"timeout" default:"60s"  env:"TIMEOUT" description:"specifies a time limit for requests"`
	Dump                bool          `long:"dump" env:"DUMP"  description:"dump or not the requests and responses while loading"`
}

type HTTPConfig struct {
	Timeout struct {
		Idle       time.Duration `long:"idle" env:"IDLE" description:"the maximum amount oftime to wait for the next request when keep-alives are enabled."`
		Read       time.Duration `long:"read" env:"READ" description:"the maximum duration for reading the entire request, including the body"`
		Write      time.Duration `long:"write" env:"ENV" description:"the maximum duration before timing out writes of the response."`
		MustShutIn time.Duration `long:"shut" env:"SHUT" default:"30s" description:"the maximum duration before timing out the graceful shutdown"`
	} `group:"timeout" namespace:"timeout" env-namespace:"TIMEOUT"`

	Host string `long:"host" env:"HOST" default:"0.0.0.0" description:"host to listen to"`
	Port int    `long:"port" env:"PORT" default:"8080" description:"port to listen to"`
	TLS  struct {
		Cert string `long:"cert" env:"CERT" description:"cert file"`
		Key  string `long:"key" env:"KEY"  description:"key file"`
	} `group:"tls opts" namespace:"tls" env-namespace:"TLS"`
}
