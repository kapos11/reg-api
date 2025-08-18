package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"reg-api/config"
	"reg-api/models"
)

func UserExists(email string) bool {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email=?"
	err := config.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		log.Printf("Error checking user: %v", err)
		return false
	}
	return count > 0
}

func RegisterUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
	}
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err = config.DB.Exec(query, user.Name, user.Email, string(hashedPassword))
	return err
}
