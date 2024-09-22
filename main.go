package main

import (
	"fmt"
	"net/http"

	"github.com/doncuervoc/go-api-02/db_connection"
	"github.com/doncuervoc/go-api-02/models"
	"github.com/doncuervoc/go-api-02/routes" // package router
	"github.com/gorilla/mux"                 // Gorilla mux = flexible package to make router HTTP
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func main() {
	fmt.Println("Initialazing Server")

	// DB Connection POSGRESQL
	// db_connection.InitDatabase() //*sql.DB
	db_connection.DBConnection() //*gorm.DB

	//Migration
	// estamos importanto el strack, creando las tablas a partir de los modelos
	db_connection.DB.AutoMigrate(models.Task{})
	db_connection.DB.AutoMigrate(models.User{})

	//ROUTER
	router := mux.NewRouter()
	//ENDPOINTS
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserById).Methods("GET")
	router.HandleFunc("/users", routes.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	fmt.Println("Ready for some margaritas!")
	http.ListenAndServe(":3000", router)

}
