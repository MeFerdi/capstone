package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Define environment variables
var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
)

var DB *sqlx.DB

// LoadConfig checks if all necessary environment variables are set
func LoadConfig() {
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatalf("Missing environment variables")
	}
	// Print environment variables for debugging purposes
	log.Printf("DB_USER: %s", dbUser)
	log.Printf("DB_PASSWORD: %s", dbPassword)
	log.Printf("DB_NAME: %s", dbName)
	log.Printf("DB_HOST: %s", dbHost)
	log.Printf("DB_PORT: %s", dbPort)
}

// DBConfig returns the connection string for the database
func DBConfig() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)
}

// InitDB initializes and returns the database connection
func InitDB() *sqlx.DB {
	connStr := DBConfig()

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	return db
}
