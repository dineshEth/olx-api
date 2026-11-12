package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func ConnectDB(dburl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dburl)
	if err != nil {
		return nil, fmt.Errorf("open database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db.PingContext(ctx)

	return db, err
}
