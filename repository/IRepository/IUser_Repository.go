package Repository

import "github.com/doncuervoc/go-api-02/models"

type IUsers_Repository interface {
	GetUsers() ([]models.User, error)        // return all users list
	GetUserById(id int) (models.User, error) // return single user by id
	CreateUser(user models.User) error       // create new user
	UpdateUser(user models.User) error       // update an existent user
	DeleteUser(id int) error                 // delete an existent user
}
