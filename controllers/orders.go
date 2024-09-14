package controllers

import (
	"capstone/config"
	"capstone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	err := config.DB.Get(&product, "SELECT * FROM products WHERE id=$1", order.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product not found"})
		return
	}
	order.TotalPrice = float64(order.Quantity) * product.Price
	order.Status = "pending"

	_, err = config.DB.NamedExec(`INSERT INTO orders (product_id, retailer_id, quantity, total_price, status) VALUES (:product_id, :retailer_id, :quantity, :total_price, :status)`, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order placed successfully!"})
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	err := config.DB.Get(&order, "SELECT * FROM orders WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
