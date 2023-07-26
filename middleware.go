package main

import (
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Set allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Set allowed methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// Allow credentials (if needed)
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests (OPTIONS method)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
