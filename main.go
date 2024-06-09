package main

import (
    "log"
    "twitter-like-backend/config"
    "twitter-like-backend/database"
    "twitter-like-backend/routers"

    "github.com/gofiber/fiber/v2"
)

func main() {
    // Load configuration
    if err := config.LoadConfig(); err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Connect to database
    database.Connect()

    // Create a new Fiber instance
    app := fiber.New()

    // Setup routes
    routers.SetupRoutes(app)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
