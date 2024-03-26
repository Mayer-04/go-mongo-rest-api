package models

// User represents a user in the database.
type User struct {

	// Name is the name of the user.
	Name string `json:"name,omitempty" bson:"name,omitempty"`

	// Email is the email of the user.
	Email string `json:"email,omitempty" bson:"email,omitempty"`

	// Password is the password of the user.
	Password string `json:"password,omitempty" bson:"password,omitempty"`

	// Age is the age of the user.
	Age uint8 `json:"age,omitempty" bson:"age,omitempty"`
}
