package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dineshEth/olx-api/internal/config"
)

func main() {
	cfg := config.MostLoad()
	mux := http.NewServeMux()

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
