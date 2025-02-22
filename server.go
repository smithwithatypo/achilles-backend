package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register handler
	http.HandleFunc("/", helloHandler)

	// Define server address
	port := ":8080"
	fmt.Println("Starting server on port", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}