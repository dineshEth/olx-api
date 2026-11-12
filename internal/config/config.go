package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Env          string
	Database_URL string
}

func MostLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("config: port is required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Fatalf("config: env is required")
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		log.Fatalf("config: database url is required")
	}

	return Config{
		Port:         port,
		Env:          env,
		Database_URL: database_url,
	}
}
