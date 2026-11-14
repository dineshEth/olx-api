package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/dineshEth/olx-api/internal/middlewares"
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

func (lh *listingHandler) Listings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	ctx := r.Context()
	requesId := ctx.Value(middlewares.RequestIdKey).(string)
	rows, err := lh.db.QueryContext(ctx, `SELECT id, title, description, price, city, created_at
				FROM listings
				ORDER BY created_at DESC
				LIMIT 100`)
	if err != nil {
		lh.logger.Error("failed query:", "request_id", requesId, "err", err)
		http.Error(w, "internal_error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	var listings = []listing{}
	for rows.Next() {
		var l listing
		if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
			lh.logger.Error("failed rows scan:", "request_id", requesId, "err", err)
			http.Error(w, "internal_error", http.StatusInternalServerError)
			return
		}
		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		lh.logger.Error("failed rows error", "request_id", requesId, "err", err)
		http.Error(w, "internal_error", http.StatusInternalServerError)
		return
	}

	lh.logger.Info("listings successfull", "request_id", requesId)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(listings); err != nil {
		lh.logger.Error("failed json encode", "request_id", requesId, "err", err)
	}
}
