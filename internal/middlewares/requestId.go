package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextkey string

const RequestIdKey contextkey = "request_id"

func generateUUID() string {
	uuid := uuid.NewString()
	return uuid
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := generateUUID()
		ctx := context.WithValue(r.Context(), RequestIdKey, requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
