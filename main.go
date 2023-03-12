package main

import (
	"log"
	"shieldcam/db"
	"shieldcam/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Establece la conexión a la base de datos
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migra los modelos para crear las tablas en la base de datos
	err = db.AutoMigrate(&models.User{}, &models.Radio{})
	if err != nil {
		log.Fatal(err)
	}

	// Crea un enrutador Gin
	r := gin.Default()

	// Define las rutas de la API
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "¡Hoxxla, mundo!"})
	})

	// Inicia el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
