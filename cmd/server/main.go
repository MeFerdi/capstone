package main

import (
	"capstone/config"
	"capstone/internal/handler"
	"capstone/internal/repository"
	"capstone/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	db, err := gorm.Open(postgres.Open(config.DBConfig()), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/products", productHandler.GetProducts)
	r.POST("/products", productHandler.CreateProduct)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
