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

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Application/Content-type", "json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
