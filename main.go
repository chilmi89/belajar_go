package main

import (
	"belajar_golang/config"
	"belajar_golang/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.GET("/users", controllers.AmbilDataUser)


	r.Run("localhost:8080")
}
