package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file. Proceeding without it: %v", err)
	}
}
