package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Register handler
	http.HandleFunc("/", helloHandler)

	// Define server address
	port := ":" + os.Getenv("PORT")
	fmt.Println("Starting server on port", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
