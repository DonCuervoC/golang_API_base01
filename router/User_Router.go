package myRouter

import (
	Controller "github.com/doncuervoc/go-api-02/controller"
	"github.com/gin-gonic/gin"
)

// UserRouter sets up the routes for user-related operations.
// It configures the HTTP methods and the corresponding controller functions.
func UserRouter(r *gin.Engine) {
	r.GET("/users", Controller.GetUsers)          // GET request to fetch all users.
	r.GET("/users/:id", Controller.GetUserById)   // GET request to fetch a user by ID.
	r.POST("/users", Controller.CreateUser)       // POST request to create a new user, with the user data in the request body.
	r.PATCH("/users/:id", Controller.UpdateUser)  // PATCH request to update a user's details, where ':id' is a path parameter.
	r.DELETE("/users/:id", Controller.DeleteUser) // DELETE request to remove a user by ID, where ':id' is a path parameter.
}

/*
The UserRouter function serves as a central point for defining routes related to user management. Each route maps an HTTP method and a URL path to a specific controller function that handles the business logic associated with that route.

This modular approach allows for clear separation of concerns within the application, enhancing maintainability and readability. Additionally, it simplifies the addition of new routes and functionality in the future.
*/
