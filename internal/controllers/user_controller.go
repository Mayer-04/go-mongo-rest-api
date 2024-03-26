package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mayer-04/go-mongo-rest-api/internal/models"
	"github.com/Mayer-04/go-mongo-rest-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
)

// User represents a user in the database.
var user models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	mensaje := "Obtener todos los usuarios"
	_, err := w.Write([]byte(mensaje))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request path
	userID := r.PathValue("id")

	// Find the user in the database
	filter := bson.D{{Key: "_id", Value: userID}}

	// Get the collection
	collection := database.GetUsersCollection()

	// Find the user
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	// Check if the user was found
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Set the response content type
	w.Header().Set("Content-Type", "application/json")

	// Write the user object to the response body
	json.NewEncoder(w).Encode(user)

	// Set the response status code
	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusCreated)
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
