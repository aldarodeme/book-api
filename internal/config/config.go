package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type Config struct {
	MySQL *MySQL
}

type MySQL struct {
	DatabaseUrl string `env:"DATABASE_URL"`
}

func LoadConfig(envLocation string) (*Config, error) {
	err := godotenv.Load(envLocation)
	if err != nil {
		log.Println("could not load .env, skipping")
	}

	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
