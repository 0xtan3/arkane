package main

import (
	"log"

	"github.com/0xtan3/arkane/internal/config"
	"github.com/0xtan3/arkane/internal/db"
)


func main() {

	cfg := config.LoadConfig()

	// Load environment variables
	// githubToken := os.Getenv("GITHUB_TOKEN")
	// domain := os.Getenv("DOMAIN")

	// Connect to MongoDB
	database, err := db.MongoConnect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
}