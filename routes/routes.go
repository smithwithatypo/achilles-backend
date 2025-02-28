package routes

import (
	"net/http"

	"github.com/smithwithatypo/achilles-backend/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/user", handlers.GetUserHandler)
	http.HandleFunc("/transcribe", handlers.TranscribeAudioHandler)

}
