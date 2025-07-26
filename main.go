package main

import (
	"log"
	"os"
	"test/middlewares"
	"test/models"
	"test/practicego"
	"test/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/pkg/config"
)

func main() {
	// Load environment config
	config.LoadConfig()

	// Get DB connection (example)
	models.ConnectDB()

	// prperties
	gin.DisableConsoleColor()
	gin.ForceConsoleColor()
	gin.ErrorLogger()

	// Create Gin router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// middlewares
	router.Use(middlewares.ThrottleGuard())

	// Use separated routes
	routes.SetupRoutes(router)

	// Start server with port from environment or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Gin server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}

	practicego.Add(2, 4)
}
