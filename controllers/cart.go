package controllers

import (
	"capstone/config"
	"capstone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.NamedExec(`INSERT INTO cart_items (retailer_id, product_id, quantity) VALUES (:retailer_id, :product_id, :quantity)`, &cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added to cart successfully!"})
}

func ViewCart(c *gin.Context) {
	retailerID := c.Param("retailer_id")
	var cartItems []models.CartItem

	err := config.DB.Select(&cartItems, "SELECT * FROM cart_items WHERE retailer_id=$1", retailerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}
