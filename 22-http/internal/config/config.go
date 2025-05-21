package config

type Config struct {
	Host   string
	Port   int
	Secret string
}

var Cfg Config

func NewConfig() *Config {
	Cfg := Config{
		Host:   "0.0.0.0",
		Port:   8989,
		Secret: "sadffs87fiofnsdkjdndsфаыпдлат",
	}
	return &Cfg
}
