package handlers

import (
	"database/sql"
	"net/http"
)

type listingHandler struct {
	db *sql.DB
}

func NewListingHandler(db *sql.DB) *listingHandler {
	return &listingHandler{
		db: db,
	}
}

func Listings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
