package main

import (
	"capstone/config"
	"capstone/models"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}

	// Store the DB connection in the config
	config.DB = db

	// Run any setup tasks here
	createTables()
	seedDatabase()
}

// createTables creates the necessary database tables
func createTables() {
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatalf("Error reading schema file: %v", err)
	}

	_, err = config.DB.Exec(string(schema))
	if err != nil {
		log.Fatalln("Failed to create tables:", err)
	}

	log.Println("Tables created successfully!")
}

// seedDatabase adds initial data to the database
func seedDatabase() {
	// Add initial users
	users := []models.User{
		{Name: "Omondi Otieno", Email: "otis@example.com", PasswordHash: "hashed_password", Role: "farmer"},
		{Name: "Mary Juma", Email: "mary@example.com", PasswordHash: "hashed_password", Role: "retailer"},
	}

	for _, user := range users {
		_, err := config.DB.NamedExec(`INSERT INTO users (name, email, password_hash, role) VALUES (:name, :email, :password_hash, :role)`, &user)
		if err != nil {
			log.Fatalln("Failed to seed users:", err)
		}
	}

	// Add initial products
	products := []models.Product{
		{UserID: 1, Name: "Tomatoes", Description: "Fresh red tomatoes", Price: 3.50, Quantity: 100},
		{UserID: 1, Name: "Cucumbers", Description: "Crisp cucumbers", Price: 2.00, Quantity: 150},
	}

	for _, product := range products {
		_, err := config.DB.NamedExec(`INSERT INTO products (user_id, name, description, price, quantity) VALUES (:user_id, :name, :description, :price, :quantity)`, &product)
		if err != nil {
			log.Fatalln("Failed to seed products:", err)
		}
	}

	log.Println("Database seeded successfully!")
}
