package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection instance.
var DB *gorm.DB

// Connection initializes the database connection.
func Connection() {
	var err error
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// Open a connection to the database.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Simple query to ensure the DB is being used correctly.
	if err := DB.Exec("SELECT 1").Error; err != nil {
		log.Fatalf("Database test query failed: %v", err)
	}

	fmt.Println("Database connection established successfully")
}
