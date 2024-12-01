package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB represents the database connection
type DB struct {
	Client *mongo.Client
}

// Connect initializes a MongoDB connection
func MongoConnect(uri string) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return &DB{Client: client}, nil
}

// Close disconnects the MongoDB client
func (db *DB) Close() error {
	return db.Client.Disconnect(context.Background())
}
