package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MostLoad() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("config: load env: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("config: port is required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Fatalf("config: env is required")
	}

	return Config{
		Port: port,
		Env:  env,
	}
}
