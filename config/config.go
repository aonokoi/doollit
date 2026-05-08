package config

import (
	"log"
	"os"

	"proj/doollit/internal/adapter/postgres"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres postgres.Config
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
