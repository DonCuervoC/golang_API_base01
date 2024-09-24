package Repository

import (
	"testing"

	"github.com/doncuervoc/go-api-02/db_connection"
	"github.com/doncuervoc/go-api-02/models"
	"github.com/stretchr/testify/assert"
)

// Test para crear un usuario
func TestCreateUser(t *testing.T) {

	// Conectar a la base de datos de desarrollo
	db_connection.DBConnection()

	// Limpiar la base de datos antes de la prueba (opcional)
	db_connection.DB.Exec("DELETE FROM users WHERE email = ?", "testuser@example.com")

	repo := UsersRepository{}

	// Crear usuario de prueba
	user := models.User{FirstName: "Test01", LastName: "User01", Email: "testuser@example.com"}
	err := repo.CreateUser(user)

	// Verificar que no hubo errores
	assert.NoError(t, err)

	// Verificar que el usuario se creó correctamente
	var createdUser models.User
	db_connection.DB.Where("email = ?", "testuser@example.com").First(&createdUser)

	// Cambiado el valor esperado para que coincida con lo creado
	assert.Equal(t, "Test01", createdUser.FirstName)
	assert.Equal(t, "testuser@example.com", createdUser.Email)

	// Limpiar después de la prueba (opcional)
	db_connection.DB.Exec("DELETE FROM users WHERE email = ?", "testuser@example.com")
}
