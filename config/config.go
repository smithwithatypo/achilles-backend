package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Load .env file if it exists, but don't fail if it doesn't
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it. Using environment variables.")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
