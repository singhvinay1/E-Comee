package main

import (
	"fmt"
	"log"

	"ecommerce-backend/config"
	"ecommerce-backend/database"
	"ecommerce-backend/routes"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Println("Configuration loaded successfully")

	// Initialize database
	database.InitDatabase()
	log.Println("Database initialized successfully")

	// Setup routes
	r := routes.SetupRoutes()
	log.Println("Routes configured successfully")

	// Start server
	serverAddr := fmt.Sprintf("%s:%s", config.AppConfig.Server.Host, config.AppConfig.Server.Port)
	log.Printf("Starting server on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 