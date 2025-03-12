package routes

import (
	"net/http"

	"github.com/smithwithatypo/achilles-backend/handlers"
	"github.com/smithwithatypo/achilles-backend/middleware"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes() {
	// Create a new router
	mux := http.NewServeMux()

	// Add your routes here
	mux.HandleFunc("/", handlers.HelloHandler)
	mux.HandleFunc("/user", handlers.GetUserHandler)
	mux.HandleFunc("/transcribe", handlers.TranscribeAudioHandler)
	mux.HandleFunc("/sentences", handlers.HandleOpenAIRequest)

	// Apply CORS middleware
	handler := middleware.CorsMiddleware(mux)

	// Set the handler with middleware
	http.Handle("/", handler)
}
