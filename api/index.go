package api

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type of the response
	w.Header().Set("Content-Type", "application/json")
	
	// Write a JSON response
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

// Define additional handlers for other API endpoints if needed
