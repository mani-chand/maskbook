package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// This variable will hold the database instance
	db *mongo.Database
)

// ConnectDB initializes the connection to MongoDB
func ConnectDB() {
	// 1. Get the MongoDB URI from environment variables
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	// 2. Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 3. Connect to the MongoDB cluster
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	}

	// 4. Ping the database to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Could not ping MongoDB: %v\n", err)
	}

	log.Println("Connected to MongoDB!")

	// 5. Store the database instance for other packages to use
	dbName := os.Getenv("MONGODB_NAME")
	if dbName == "" {
		dbName = "noob_db" // Add a sensible default
	}
	db = client.Database(dbName)
}

// GetDatabase returns the connected database instance
func GetDatabase() *mongo.Database {
	return db
}
