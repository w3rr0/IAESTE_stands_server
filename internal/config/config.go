package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
}

var DatabaseURL = os.Getenv("DATABASE_URL")
var EmailUser = os.Getenv("EMAIL_USER")
var EmailPassword = os.Getenv("EMAIL_PASSWORD")
