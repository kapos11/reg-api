package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"reg-api/models"
	"reg-api/services"
	"reg-api/utils"
)

func RegisterHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed: " + err.Error()})
		return
	}

	if services.UserExists(user.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	if err := services.RegisterUser(user); err != nil {
		log.Printf("Insert error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
