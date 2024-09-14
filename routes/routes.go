package routes

import (
	"capstone/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// User routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Product routes
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)

	// Order routes
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.GetOrder)

	// Cart routes
	r.POST("/cart", controllers.AddToCart)
	r.GET("/cart/:retailer_id", controllers.ViewCart)
}
