package Repository

import (
	"github.com/doncuervoc/go-api-02/db_connection"
	"github.com/doncuervoc/go-api-02/models"
)

// UsersRepository struct that implements the IUsers_Repository interface.
// This struct provides methods for CRUD operations on user entities.
type UsersRepository struct{}

// GetUsers retrieves a list of all users from the database.
// It returns a slice of User models and an error if any occurred during the operation.
func (r *UsersRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := db_connection.DB.Find(&users).Error // Queries all users and stores them in 'users'.
	if err != nil {
		return nil, err // Returns nil and the error if the query fails.
	}
	return users, nil // Returns the list of users and nil if successful.
}

// GetUserById retrieves a single user from the database by its ID.
// It takes an integer 'id' as an argument and returns the corresponding User model and an error.
func (r *UsersRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := db_connection.DB.First(&user, id).Error // Fetches the first user that matches the provided ID.
	if err != nil {
		return models.User{}, err // Returns an empty User model and the error if the user is not found.
	}
	return user, nil
}

// CreateUser inserts a new user into the database.
// It takes a User model as an argument and returns an error if the operation fails.
func (r *UsersRepository) CreateUser(user models.User) error {
	return db_connection.DB.Create(&user).Error // Creates a new user record in the database.
}

// UpdateUser modifies an existing user in the database.
// It takes a User model as an argument and returns an error if the update operation fails.
func (r *UsersRepository) UpdateUser(user models.User) error {
	return db_connection.DB.Save(&user).Error // Updates the existing user record in the database.
}

// DeleteUser removes a user from the database by its ID.
// It takes an integer 'id' as an argument and returns an error if the deletion fails.
// Note: This method performs a soft delete by default.
func (r *UsersRepository) DeleteUser(id int) error {
	var user models.User
	err := db_connection.DB.First(&user, id).Error // Attempts to find the user by ID before deletion.
	if err != nil {
		return err
	}

	return db_connection.DB.Delete(&user).Error //SOFT delete
	// return db_connection.DB.Unscoped().Delete(&user).Error //HARD delete

}
