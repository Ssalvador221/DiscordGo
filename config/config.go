package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("No token provided")
	}
	return token
}
