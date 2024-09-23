package Controller

import (
	"net/http"
	"strconv"

	"github.com/doncuervoc/go-api-02/models"
	Repository "github.com/doncuervoc/go-api-02/repository"
	"github.com/gin-gonic/gin"
)

// GetUsers retrieves the list of all users.
// It interacts with the UsersRepository to fetch user data and handles HTTP responses accordingly.
func GetUsers(c *gin.Context) {
	usersRepo := Repository.UsersRepository{} // Create a new instance of UsersRepository.
	users, err := usersRepo.GetUsers()        // Fetch all users from the repository.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users: " + err.Error()}) // Return a 500 error response if fetching fails.
		return
	}
	c.JSON(http.StatusOK, users) // Return a 200 response with the list of users if successful.
}

// GetUserById retrieves a user by their ID.
// It takes the ID from the request parameters, converts it to an integer, and fetches the user from the repository.
func GetUserById(c *gin.Context) {
	usersRepo := Repository.UsersRepository{} // Create a new instance of UsersRepository.
	id, err := strconv.Atoi(c.Param("id"))    // Convert the 'id' parameter from string to int.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Return a 400 error response if the ID is invalid.
		return
	}

	user, err := usersRepo.GetUserById(id) // Fetch the user by ID from the repository.
	if err != nil || user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // Return a 404 error response if the user is not found.
		return
	}
	c.JSON(http.StatusOK, user) // Return a 200 response with the found user if successful.
}

// CreateUser creates a new user based on the input provided in the request body.
// It binds the JSON input to a User model and uses the UsersRepository to insert it into the database.
func CreateUser(c *gin.Context) {
	usersRepo := Repository.UsersRepository{} // Create a new instance of UsersRepository.
	var user models.User                      // Declare a variable to hold the new user data.

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()}) // Return a 400 error response if input is invalid.
		return
	}

	if err := usersRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user: " + err.Error()}) // Return a 500 error response if user creation fails
		return
	}

	c.JSON(http.StatusCreated, user) // Return a 201 response with the created user if successful.
}

// UpdateUser updates an existing user based on the provided ID and user data in the request body.
// It retrieves the user, validates the input, and applies the changes via the UsersRepository.
func UpdateUser(c *gin.Context) {
	usersRepo := Repository.UsersRepository{} // Create a new instance of UsersRepository.
	id, err := strconv.Atoi(c.Param("id"))    // Convert the 'id' parameter from string to int.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Return a 400 error response if the ID is invalid.
		return
	}

	user, err := usersRepo.GetUserById(id) // Fetch the existing user by ID from the repository.
	if err != nil || user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // Return a 404 error response if the user is not found.
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()}) // Return a 400 error response if input is invalid.
		return
	}

	if err := usersRepo.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user: " + err.Error()}) // Return a 500 error response if the update fails
		return
	}

	c.JSON(http.StatusOK, user) // Return a 200 response with the updated user if successful.
}

// DeleteUser removes a user by their ID from the database.
// It verifies the ID, attempts to delete the user using the UsersRepository, and handles the HTTP response.
func DeleteUser(c *gin.Context) {
	usersRepo := Repository.UsersRepository{} // Create a new instance of UsersRepository.
	id, err := strconv.Atoi(c.Param("id"))    // Convert the 'id' parameter from string to int.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"}) // Return a 400 error response if the ID is invalid.
		return
	}

	if err := usersRepo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user: " + err.Error()}) // Return a 500 error response if the deletion fails.
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"}) // Return a 200 response indicating successful deletion.
}

/*
Instance of the repository: In each controller function, I create an instance of UsersRepository that implements the IUsers_Repository interface. Through this instance, we call the repository methods.

Separation of database logic from the controller: Now, database operations are encapsulated within the repository, allowing the controller to focus only on handling HTTP logic.

Scalability improvements: This design enables future changes to the repository implementation without requiring modifications to the controller, promoting better maintainability and testability.
*/
