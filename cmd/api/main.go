package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dineshEth/olx-api/internal/config"
	"github.com/dineshEth/olx-api/internal/db"
	"github.com/dineshEth/olx-api/internal/handlers"
)

func main() {
	cfg := config.MostLoad()
	mux := http.NewServeMux()
	_, err := db.ConnectDB(cfg.Database_URL)
	if err != nil {
		log.Fatalf("database connection fail %v", err)
	}

	mux.HandleFunc("GET /", handlers.HomeHandler)
	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server startup fail %v", err)
	}
}
