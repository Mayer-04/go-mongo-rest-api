package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mayer-04/go-mongo-rest-api/internal/models"
	"github.com/Mayer-04/go-mongo-rest-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the database.
var user models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {

	collection := database.GetUsersCollection()

	var users []models.User

	// Find all users - bson.M{} is an empty filter that returns all documents
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// Close the cursor when the function returns
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// Set the response content type
	w.Header().Set("Content-Type", "application/json")

	// Set the response status code
	w.WriteHeader(http.StatusOK)

	// Write the user object to the response body
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request path
	userID := r.PathValue("id")

	// Get the collection
	collection := database.GetUsersCollection()

	// Convert the user ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find the user in the database - bson.D{{Key: "_id", Value: userID}}
	filter := bson.M{"_id": objectID}

	// Find the user
	err = collection.FindOne(context.Background(), filter).Decode(&user)

	// Check if the user was found
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Set the response content type
	w.Header().Set("Content-Type", "application/json")

	// Set the response status code
	w.WriteHeader(http.StatusOK)

	// Write the user object to the response body
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	collection := database.GetUsersCollection()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	collection := database.GetUsersCollection()

	userID := r.PathValue("id")

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": userID})
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
