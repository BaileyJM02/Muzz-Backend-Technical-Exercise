package server

import (
	"fmt"
	"net/http"
	"time"
)

// Start initiates the HTTP server
func Start() {
	// Start the HTTP server ()
	server := &http.Server{
		Addr:         ":3000",
		Handler:      GetRouter(),
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
	}

	// Inform us that we are starting the server
	fmt.Printf("Starting server on %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %v", err)
	}
}
