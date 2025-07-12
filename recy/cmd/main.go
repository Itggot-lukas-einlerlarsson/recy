package main

import (
    "log"
    "net/http"
    "recy/internal/handlers"
    "recy/pkg/config"
    "recy/pkg/middleware"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize handlers
    handlers.InitHandlers()

    // Set up middleware
    mux := middleware.SetupMiddleware(http.NewServeMux())

    // Register routes
    mux.HandleFunc("/", handlers.HomeHandler)

    // Start server
    log.Printf("Server starting on %s", cfg.ServerAddress)
    log.Fatal(http.ListenAndServe(cfg.ServerAddress, mux))
}
