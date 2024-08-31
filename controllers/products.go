package controllers

import (
	"capstone/config"
	"capstone/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := []models.Product{}
	err := config.DB.Select(&products, "SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func SearchAndFilterProducts(c *gin.Context) {
	var products []models.Product
	searchQuery := c.Query("q")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")

	query := "SELECT * FROM products WHERE 1=1"

	if searchQuery != "" {
		query += " AND (name ILIKE '%' || $1 || '%' OR description ILIKE '%' || $1 || '%')"
	}

	if minPrice != "" {
		query += " AND price >= $2"
	}

	if maxPrice != "" {
		query += " AND price <= $30"
	}

	err := config.DB.Select(&products, query, searchQuery, minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.NamedExec(`INSERT INTO products (user_id, name, description, price, quantity) VALUES (:user_id, :name, :description, :price, :quantity)`, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully!"})
}
