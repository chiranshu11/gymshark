package config

import (
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("PORT"),
	}
	return cfg, nil
}
