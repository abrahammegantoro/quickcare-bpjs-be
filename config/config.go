package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBUrl string `env:"DB_URL,required"`
	CollectionName string `env:"COLLECTION_NAME,required"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: %e", err)
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatal("Error parsing env config: %e", err)
	}

	return config
}