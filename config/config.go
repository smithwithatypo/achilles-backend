package config

import (
	// "log"
	"os"

	// "github.com/joho/godotenv"
)

// func LoadConfig() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func GetEnv(key string) string {
	return os.Getenv(key)
}
