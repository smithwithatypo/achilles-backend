package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smithwithatypo/achilles-backend/config"
	_ "github.com/smithwithatypo/achilles-backend/middleware" // Import middleware package
	"github.com/smithwithatypo/achilles-backend/routes"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Register routes
	routes.RegisterRoutes()

	// Define server address
	port := ":" + config.GetEnv("PORT")
	fmt.Println("Starting server on port", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
