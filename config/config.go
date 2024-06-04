package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	UserServiceURL    string
	ProductServiceURL string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	UserServiceURL = os.Getenv("USER_SERVICE_URL")
	ProductServiceURL = os.Getenv("PRODUCT_SERVICE_URL")
}
