// database.go
package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Obtiene las credenciales de la base de datos desde variables de entorno
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")
	dbName := os.Getenv("PGDATABASE")

	// Establece la cadena de conexión a la base de datos
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	// Establece la conexión a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
