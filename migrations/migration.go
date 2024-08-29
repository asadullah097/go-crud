package main

import (
	"fmt"
	"go-curd/initializers"
	"go-curd/models"
	"log"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.Connection()
}

func main(){
	
	if initializers.DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	// Migrate the Post model
	if err := initializers.DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	fmt.Println("Migration completed successfully")
}