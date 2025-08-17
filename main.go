package main

import (
	"database/sql"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var validate *validator.Validate

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func main() {
	initDB()
	defer db.Close()

	validate = validator.New()
	router := gin.Default()

	router.POST("/register", registerHandler)

	log.Println("Server running on :8080")
	log.Fatal(router.Run(":8080"))
}

func initDB() {
	var err error
	dsn := "myuser:mypassword@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB connect failed: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Create table failed: %v", err)
	}
}

func registerHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed: " + err.Error()})
		return
	}

	if userExists(user.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	if err := registerUser(user); err != nil {
		log.Printf("Insert error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func userExists(email string) bool {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email=?"
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		log.Printf("Error checking user: %v", err)
		return false
	}
	return count > 0
}

func registerUser(user User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
	}
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err = db.Exec(query, user.Name, user.Email, string(hashedPassword))
	return err
}
