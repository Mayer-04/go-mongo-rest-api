package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mayer-04/go-mongo-rest-api/internal/models"
	"github.com/Mayer-04/go-mongo-rest-api/pkg/database"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the database.
var user models.User

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    models.User `json:"data,omitempty"`
}

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

	// Iterate through the cursor and decode each user into the user variable
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
		http.Error(w, "invalid user ID", http.StatusBadRequest)
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
	// Get user collection from database
	collection := database.GetUsersCollection()

	// Create a new "validator" instance
	validate := validator.New()

	// Try to decode the request body in the "user" variable
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate user data
	if err := validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user into the collection
	// InsertOne returns an "InsertOneResult" with the inserted ID of the new document
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sets the content type of the response
	w.Header().Set("Content-Type", "application/json")

	// Set response status code
	w.WriteHeader(http.StatusCreated)

	// Encode the user in the response body
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	collection := database.GetUsersCollection()
	validator := validator.New()

	userID, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	// Create a filter to search for the user by their ID
	filter := bson.M{"_id": userID}

	// Create a new user with updated data
	var userUpdate models.User

	// Decode the request body in the new user
	err = json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user data
	if err := validator.Struct(userUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new filter to update the user in the database
	update := bson.M{"$set": userUpdate}

	// "UpdateResult" contains the number of updated documents and an error
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the user was updated in the database
	// Number of documents that match the filter, if it is 0 it is not updated
	if result.ModifiedCount == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userUpdate)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	collection := database.GetUsersCollection()

	userID := r.PathValue("id")

	objectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": objectID}

	// DeleteOne returns a "DeleteResult" with the number of documents deleted and an error
	// If there are no matches to the filter, no document is deleted and "DeleteResult" is set to 0.
	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: true,
		Message: "user deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
