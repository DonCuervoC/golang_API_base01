package main

import (
	"fmt"

	"github.com/doncuervoc/go-api-02/db_connection"
	"github.com/doncuervoc/go-api-02/models"
	myRouter "github.com/doncuervoc/go-api-02/router"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func main() {
	fmt.Println("Initialazing Server")

	// DB Connection POSGRESQL
	db_connection.DBConnection() //*gorm.DB

	// Migration: Automatically create tables based on the defined models.
	// This ensures that the database schema is in sync with the application's models
	db_connection.DB.AutoMigrate(models.User{})

	//ROUTER : Setting up the Gin router.
	router := gin.Default() // Initialize a new Gin router with default middleware.

	// ROUTES :  Configure user-related routes by calling the UserRouter function.
	myRouter.UserRouter(router) // Registering user routes.

	// Start the Gin server on port 3000 and log any errors encountered during startup.
	if err := router.Run(":3000"); err != nil {
		fmt.Println("Error starting server:", err) // Log the error if server fails to start.
	}

}
