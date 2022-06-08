package config

import (
	"github.com/caarlos0/env/v6"
)

// Config provides configuration settings
type Config struct {
	GRPCPort   string `env:"GRPC_PORT,notEmpty"`
	DBUsername string `env:"DB_USERNAME,notEmpty"`
	DBPassword string `env:"DB_PASSWORD,notEmpty"`
	DBHost     string `env:"DB_HOST,notEmpty"`
	DBPort     string `env:"DB_PORT,notEmpty"`
	DBName     string `env:"DB_NAME,notEmpty"`
}

// LoadConfig load config from environment variables
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
