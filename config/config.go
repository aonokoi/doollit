package config

import (
	"log"
	"os"

	"proj/doollit/internal/adapter/postgres"
	"proj/doollit/pkg/httpserver"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres postgres.Config
	HTTP     httpserver.Config
}

func InitConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Unable to load config: %s", err)
		os.Exit(1)
	}

	return &cfg, nil
}
