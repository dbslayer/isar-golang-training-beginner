package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configuration := Configuration{}
	configuration.DB_USERNAME = os.Getenv("DB_USERNAME")
	configuration.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	configuration.DB_PORT = os.Getenv("DB_PORT")
	configuration.DB_HOST = os.Getenv("DB_HOST")
	configuration.DB_NAME = os.Getenv("DB_NAME")
	return configuration
}
