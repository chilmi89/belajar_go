package routes


import (
	"belajar_golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", controllers.AmbilDataUser)
		
		
	}
}
