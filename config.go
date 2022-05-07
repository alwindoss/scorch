package scorch

type Config struct {
	Port       int  `env:"SCORCH_PORT" envDefault:"8080"`
	Production bool `env:"SCORCH_PRODUCTION"`
}
