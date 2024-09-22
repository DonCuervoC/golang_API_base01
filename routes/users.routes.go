package routes

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/doncuervoc/go-api-02/db_connection"
	"github.com/doncuervoc/go-api-02/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Get Users"))
	var users []models.User

	db_connection.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	var user models.User
	params := mux.Vars(r)
	// fmt.Println(params)
	// fmt.Println(params["id"])

	db_connection.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	// asignat los datos del request a la variable user, utilizamos un puntero
	json.NewDecoder((r.Body)).Decode(&user)

	newUser := db_connection.DB.Create(&user)
	err := newUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}
	// w.Write([]byte("Create User"))
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	// Buscar el usuario por ID
	db_connection.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// Decodificar el body JSON en el modelo User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input: " + err.Error()))
		return
	}

	// Actualizar el usuario en la base de datos
	db_connection.DB.Save(&user)

	// Responder con el usuario actualizado
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db_connection.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db_connection.DB.Delete(&user) // unicamente agregar delatetime into DB
	//db_connection.DB.Unscoped().Delete(&user) // Eliminar definitivamente de la DB
	w.WriteHeader(http.StatusOK)

}
