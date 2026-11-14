package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type listingHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewListingHandler(db *sql.DB, logger *slog.Logger) *listingHandler {
	return &listingHandler{
		db:     db,
		logger: logger,
	}
}

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ls *listingHandler) Listings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	ctx := r.Context()
	rows, err := ls.db.QueryContext(ctx, `SELECT id, title, description, price, city, created_at
				FROM listings
				ORDER BY created_at DESC
				LIMIT 100`)
	if err != nil {
		http.Error(w, "internal_error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	var listings = []listing{}
	for rows.Next() {
		var l listing
		if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
			http.Error(w, "internal_error", http.StatusInternalServerError)
			return
		}
		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "internal_error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(listings); err != nil {
		log.Printf("encode listings: %v", err)
	}
}
