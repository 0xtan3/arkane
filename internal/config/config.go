package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
}

func LoadConfig() Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get Mongo URI from environment variable
	URI := os.Getenv("MONGO_URI")
	if URI == "" {
		log.Fatalf("MONGO_URI environment variable not set")
	}

	return Config{
		MongoURI: URI,
	}
}
