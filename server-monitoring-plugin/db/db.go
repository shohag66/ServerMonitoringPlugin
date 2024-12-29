package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/servermonitoring_plugin"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	query := `
	CREATE TABLE IF NOT EXISTS configurations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		threshold_percentage FLOAT NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := DB.Exec(query); err != nil {
		log.Fatalf("Failed to create configurations table: %v", err)
	}

	fmt.Println("Database connected and table initialized.")
}
