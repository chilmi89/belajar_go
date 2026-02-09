package main

import (
	"belajar_golang/config"
	"belajar_golang/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.GET("/users", controllers.AmbilDataUser)

	// Ambil port dari environment variable, default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port) // bind ke semua interface
}
