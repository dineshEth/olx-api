package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB(dburl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dburl)
	if err != nil {
		return nil, fmt.Errorf("open database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %v", err)
	}

	log.Printf("Connected to database: %s", dburl)

	return db, err
}
