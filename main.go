package main

import (
	"fmt"
	"log"
	"net"
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

	// // Define server address
	// port := ":" + config.GetEnv("PORT")
	// fmt.Println("Starting server on port", port)

	// // Start the server
	// if err := http.ListenAndServe(port, nil); err != nil {
	// 	log.Fatal("Server failed:", err)
	// }

	    // Define server address
		port := config.GetEnv("PORT")
		if port == "" {
			port = "8080" // Default port if not specified
		}
		
		// Create a listener that accepts both IPv4 and IPv6 connections
		listener, err := net.Listen("tcp", ":"+port)
		if err != nil {
			log.Fatal("Failed to create listener:", err)
		}
		
		fmt.Println("Starting server on port", port)
		fmt.Println("Listening on IPv4 and IPv6")
	
		// Start the server with the dual-stack listener
		if err := http.Serve(listener, nil); err != nil {
			log.Fatal("Server failed:", err)
		}
}
