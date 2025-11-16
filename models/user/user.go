package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthInput defines a struct to bind JSON data from requests.
// ... (this struct is fine)
type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile"   binding:"required"`
}

// User defines the model for a user, updated for MongoDB.
type User struct {
	// 1. Use primitive.ObjectID for MongoDB's native _id
	ID       primitive.ObjectID `json:"id"       bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email"    bson:"email"`
	// 2. Make sure this field is for the hash, not plain text
	PasswordHash string `json:"-"        bson:"password_hash"` // Never send password to client
	Mobile       string `json:"mobile"   bson:"mobile"`
	// 3. Added the new Avatar field with corrected tags
	Avatar string `json:"avatar"   bson:"avatar,omitempty"`
}
