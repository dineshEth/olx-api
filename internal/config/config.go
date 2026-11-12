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
	godotenv.Load()

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
