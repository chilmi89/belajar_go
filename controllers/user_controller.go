package controllers

import (
	"net/http"

	"belajar_golang/config"
	"belajar_golang/models"

	"github.com/gin-gonic/gin"
)

func AmbilDataUser(c *gin.Context){
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}