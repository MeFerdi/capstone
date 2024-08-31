package main

import (
	"capstone/config"
	"capstone/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Load configurations
	config.LoadConfig()

	// Initialize database
	db := config.InitDB()
	defer db.Close()

	// Set up routes
	routes.SetupRoutes(r)

	r.Run(":8080")

}
