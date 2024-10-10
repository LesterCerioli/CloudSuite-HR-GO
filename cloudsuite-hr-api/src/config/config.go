package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))

	var err error
	maxAttempts := 5
	for attempts := 0; attempts < maxAttempts; attempts++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Printf("Database connection established successfully.")
			return
		}

		log.Printf("Failed to connect to database, retrying in 5 seconds... (%d/%d)\n", attempts+1, maxAttempts)
		time.Sleep(5 * time.Second)
	}

	panic(fmt.Sprintf("Failed to connect to database after %d attempts: %v", maxAttempts, err))

}
