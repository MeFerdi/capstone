package main

import (
	"capstone/config"
	"capstone/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load configurations
	config.LoadConfig()

	// Initialize database
	db := config.InitDB()
	defer db.Close()

	// Set up routes
	routes.SetupRoutes(r)

	r.Run(":8080")
}
