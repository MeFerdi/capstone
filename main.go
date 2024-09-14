package main

import (
	"capstone/config"
	"capstone/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConfig := config.NewDatabaseConfig(
		"postgres",    // User
		"ferdinand",   // Password
		"capstone_db", // Database name
		"localhost",   // Host
		"5432",        // Port
	)

	// Initialize database
	db := config.InitDB(dbConfig)
	defer db.Close()

	// Set up routes
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
