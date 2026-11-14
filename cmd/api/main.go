package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/dineshEth/olx-api/internal/config"
	"github.com/dineshEth/olx-api/internal/db"
	"github.com/dineshEth/olx-api/internal/handlers"
	"github.com/dineshEth/olx-api/internal/middlewares"
)

func main() {
	cfg := config.MostLoad()
	db, err := db.ConnectDB(cfg.Database_URL)
	if err != nil {
		log.Fatalf("database connection fail %v", err)
	}

	jsonLogHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	})
	jl := slog.New(jsonLogHandler)
	lh := handlers.NewListingHandler(db, jl)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.HomeHandler)
	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)
	mux.HandleFunc("GET /listings", lh.Listings)
	handler := middlewares.RequestID(mux)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("server running on :", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server startup fail %v", err)
	}
}
