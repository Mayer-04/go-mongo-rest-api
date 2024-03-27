package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	envMongoURI = "MONGODB_URI"
	database    = "go_crud"
	collection  = "users"
)

func ConnectToMongoDB() *mongo.Client {

	// Check if the environment variable is set
	var uri string = os.Getenv(envMongoURI)
	if uri == "" {
		log.Fatal("you must set your 'MONGODB_URI' environment variable")
	}

	// MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a new MongoDB client and connect
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB ")

	return client
}

func GetUsersCollection() *mongo.Collection {
	// Get the database and collection
	client := ConnectToMongoDB()
	db := client.Database(database)
	collection := db.Collection(collection)

	return collection
}
