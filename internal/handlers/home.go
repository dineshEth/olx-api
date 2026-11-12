package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Application/Content-type", "json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Welcome to the OLX API"}`))
}
