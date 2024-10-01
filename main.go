package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to represent a response
type Response struct {
	Message string `json:"message"`
}

// HelloHandler responds with a simple message
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Set up the HTTP routes
	http.HandleFunc("/hello", HelloHandler)

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
