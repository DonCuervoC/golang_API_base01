package db_connection

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Para cargar el archivo .env
	_ "github.com/lib/pq"      // Driver de PostgreSQL
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variables globales para la configuración de la base de datos
var (
	// DB       *sql.DB
	DB       *gorm.DB
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	DNS      string
	// DB     *gorm.DB
)

// var DB *gorm.DB

// Cargar las variables de entorno al iniciar el paquete
func init() {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Asignar las variables de entorno a las globales
	Host = os.Getenv("DB_HOST")
	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	Port = os.Getenv("DB_PORT")

	DNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Host, User, Password, DBName, Port)

}

/*
// InitDatabase inicializa la conexión a PostgreSQL
func InitDatabase() {
	var err error

	// Cadena de conexión para PostgreSQL
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Host, User, Password, DBName, Port)

	// Abrir la conexión con PostgreSQL
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Verificar la conexión
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("M:01 = Database connected successfully!")
}

// GetDB exporta la conexión de la base de datos
func GetDB() *sql.DB {
	return DB
}
*/

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("M:02 = Database connected successfully!")
	}
}
