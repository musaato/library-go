package initializers

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := getDBURL()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB.Exec("SET search_path TO public")

	if err != nil {
		log.Printf("Failed to connect to database. %v", err)
	}
}

func getDBURL() string {
	// Detect if running inside Docker container
	if strings.Contains(os.Getenv("DOCKER_ENV"), "true") {
		return os.Getenv("DB_URL_CONTAINER")
	}

	// Default to local development DB
	return os.Getenv("DB_URL_LOCAL")
}
