package config

import (
	"context"
	"wachirawut_agnos_backend/pkg/database/postgresql"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	PostgresqlClient *postgresql.Config
}

func NewConfig() *Config {
	var config Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		panic(err)
	}
	return &config
}
