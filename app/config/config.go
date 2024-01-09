package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server Server
}

type Server struct {
	Address string `envconfig:"ADDRESS" default:"0.0.0.0"`
	Port    string `envconfig:"PORT" default:"8080"`
}

var config Config

func Load() (*Config, error) {
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
