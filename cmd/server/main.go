package main

import (
	"fmt"
	"log"
	"net/http"

	"gymshark/internal/config"
	"gymshark/internal/handlers"
	"gymshark/internal/storage"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize in-memory storage with default pack sizes
	store := storage.NewMemoryStorage([]int{250, 500, 1000, 2000, 5000})

	// Setup HTTP handlers and routes
	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux, store)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
