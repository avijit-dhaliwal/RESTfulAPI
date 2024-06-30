package main

import (
    "log"
    "net/http"

    "github.com/avijit-dhaliwal/RESTfulAPI/internal/api"
	"github.com/avijit-dhaliwal/RESTfulAPI/internal/db"
    "github.com/avijit-dhaliwal/RESTfulAPI/internal/config"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Initialize database
    database, err := db.InitDB(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer database.Close()

    // Initialize router
    router := api.NewRouter(database)

    // Start server
    log.Printf("Server starting on port %s", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}