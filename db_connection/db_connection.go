package db_connection

import (
	"fmt"
	"log"
	"os" //For working with environment variables.

	"github.com/joho/godotenv" // For loading environment variables from a .env file.
	_ "github.com/lib/pq"      // PostgreSQL driver for GORM.
	"gorm.io/driver/postgres"  // GORM's PostgreSQL driver.
	"gorm.io/gorm"             // Importing GORM for ORM functionality.
)

// Global variables for database string connection configuration
var (
	DB       *gorm.DB // Database connection instance
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	DNS      string // Data Source Name for connecting to the database(string connection)
)

// init function is called when the package is initialized.
// It loads environment variables and configures the database connection settings.
func init() {
	// Load the .env file to access environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Assign environment variables to the global variables for database configuration
	Host = os.Getenv("DB_HOST")
	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	Port = os.Getenv("DB_PORT")
	// Construct the Data Source Name (DSN) for connecting to PostgreSQL
	DNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Host, User, Password, DBName, Port)
}

// DBConnection establishes a connection to the PostgreSQL database using GORM.
func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{}) // Open a connection to the database

	// Check for connection errors
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connected successfully!")
		log.Println("Ready for some margaritas pendejo!")
	}
}

/*
The db_connection package handles loading environment variables and establishing
a connection to the PostgreSQL database using GORM. The use of a .env file allows
for flexible configuration without hardcoding sensitive information into the source code.
*/
