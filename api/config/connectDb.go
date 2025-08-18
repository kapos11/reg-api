package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "myuser:mypassword@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB connect failed: %v", err)
	}

	if err = DB.Ping(); err != nil {
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
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Create table failed: %v", err)
	}
}
