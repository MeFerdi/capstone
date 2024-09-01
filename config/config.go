package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DatabaseConfig holds the database configuration parameters
type DatabaseConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

var DB *sqlx.DB

// NewDatabaseConfig creates a new instance of DatabaseConfig with the given values
func NewDatabaseConfig(user, password, name, host, port string) *DatabaseConfig {
	return &DatabaseConfig{
		User:     user,
		Password: password,
		Name:     name,
		Host:     host,
		Port:     port,
	}
}

// DBConfig returns the connection string for the database
func (config *DatabaseConfig) DBConfig() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Name, config.Password)
}

// InitDB initializes and returns the database connection
func InitDB(config *DatabaseConfig) *sqlx.DB {
	connStr := config.DBConfig()
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}
	DB = db
	return db
}
