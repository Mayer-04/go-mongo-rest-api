package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the database.
type User struct {

	// ID is the unique identifier of the user.
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	// Name is the name of the user.
	Name string `bson:"name,omitempty" json:"name,omitempty" validate:"required"`

	// Email is the email of the user
	Email string `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`

	// Password is the password of the user.
	Password string `bson:"password,omitempty" json:"password,omitempty" validate:"required,min=8"`

	// Age is the age of the user.
	Age uint8 `bson:"age,omitempty" json:"age,omitempty" validate:"gte=0,lte=130"`
}
