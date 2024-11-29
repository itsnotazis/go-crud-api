package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env *string
}

func EnvConfig() *Config {
	return &Config{}
}

func (c *Config) LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
}

func (c *Config) GetEnv(key string, fallback string) {
	env := os.Getenv(key)

	if env != key {
		env = fallback
	}
}
