package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"reg-api/config"
	"reg-api/handlers"
	"reg-api/utils"
)

func main() {
	// dbconnect
	config.InitDB()
	defer config.DB.Close()
	utils.InitValidator()

	// سيرفر Gin
	router := gin.Default()

	// تعريف الـ endpoints
	router.POST("/register", handlers.RegisterHandler)

	log.Println("Server running on :8081")
	log.Fatal(router.Run(":8081"))
}
