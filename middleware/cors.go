package middleware

import (
	"net/http"
	"os"
)

// CorsMiddleware adds CORS headers to all responses
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set allowed origin based on environment
        allowedOrigin := "*"
        environment := os.Getenv("ENVIRONMENT")
        
        if environment == "production" {
            allowedOrigin = "https://achilles-frontend-production.up.railway.app"
        } else if environment == "development" {
            allowedOrigin = "http://localhost:4173"
        }

		// Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
}
