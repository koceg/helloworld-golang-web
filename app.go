package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new HTTP request multiplexer
	mux := http.NewServeMux()

	// Create a new handler
	handler := new(Handler)

	// Set routes to their handlers
	mux.Handle("/", handler)

	// Create a new http.Server object, tuning its parameters.
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	// Listen to requests
	server.ListenAndServe()
}

type Handler struct{}

// A HTTP Handler must implement ServeHTTP interface
func (h *Handler) ServeHTTP(
	writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}
