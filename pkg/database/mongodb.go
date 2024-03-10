package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() *mongo.Collection {

	// Get the MongoDB URI from the environment variable
	const envMongoURI = "MONGODB_URI"

	// Check if the environment variable is set
	uri := os.Getenv(envMongoURI)
	if uri == "" {
		log.Fatal("MONGODB_URI environment variable is not set")
	}

	// Create a new MongoDB client and connect
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Disconnect from MongoDB when the program exits
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Get the database and collection
	db := client.Database("go_crud")
	collection := db.Collection("users")

	return collection
}
