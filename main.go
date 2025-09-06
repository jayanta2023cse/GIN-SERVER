package main

import (
	"app/kafka"
	"app/middlewares"
	"app/models"
	"app/routes"
	"log"
	"os"

	_ "app/docs"

	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/pkg/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load environment config
	config.LoadConfig()

	// Get DB connection (example)
	models.ConnectDB()

	// Initialize Kafka
	kafka.InitKafka()

	// prperties
	// gin.DisableConsoleColor()
	gin.ForceConsoleColor()
	gin.ErrorLogger()

	// Create Gin router
	router := gin.Default()
	router.SetTrustedProxies(nil)
	// middlewares
	router.Use(middlewares.ThrottleGuard())
	// Use separated routes
	routes.SetupRoutes(router)
	// Serve Swagger UI at /swagger/*any
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server with port from environment or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Gin server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}

	// Close Kafka resources
	// defer kafka.CloseProducer()
	// defer kafka.CloseConsumer()
}
