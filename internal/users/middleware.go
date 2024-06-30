package users

import (
	"context"
	"net/http"
	"log"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(w, "Invalid token.")
			return
		}

		ctx := context.WithValue(r.Context(), "user-id", 42)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}