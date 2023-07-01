package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// NewDatabase : intializes and returns mysql db
func ConnectToDB() *gorm.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", HOST, USER, PASS, DBNAME)
	fmt.Println(DSN)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	DB = db

	return DB
}
