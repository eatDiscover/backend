package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a route handler function for the /hello endpoint
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Write a response to the client
		fmt.Fprintf(w, "Hello, World!")
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
