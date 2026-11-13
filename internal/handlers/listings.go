package handlers

import "net/http"

func Listings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
