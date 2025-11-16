package user

import (
	"context"
	// "errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	// 1. THIS IS THE CORRECT IMPORT.
	// It must match the import in main.go
	"noob/services/database"

	// 2. Use an alias for your models
	usermodels "noob/models/user"
)

// ValidateUser checks if a user's credentials are correct.
func ValidateUser(c *gin.Context) {
	var input usermodels.AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// GetDB() will no longer return nil because main.go initialized this *exact* package
	db := database.GetDatabase()
	userCollection := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var existingUser usermodels.User
	err := userCollection.FindOne(ctx, bson.M{"username": input.Username}).Decode(&existingUser)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	existingUser.PasswordHash = ""
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": existingUser})
}

// CreateUser handles the logic for creating a new user.
func CreateUser(c *gin.Context) {
	// BUG FIX 1: Bind to AuthInput to get the plain-text password
	var input usermodels.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// --- Database Logic ---
	// GetDB() will no longer return nil
	db := database.GetDatabase()
	userCollection := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Check if user already exists
	count, err := userCollection.CountDocuments(ctx, bson.M{"username": input.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking database"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	// (You should also check for existing email)

	// 2. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 3. Create a new user object
	newUser := usermodels.User{
		ID:           primitive.NewObjectID(), // Generate a new BSON ObjectID
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword), // Use "PasswordHash", not "Password"
		Mobile:       input.Mobile,
		Avatar:       "", // Initialize new Avatar field as empty
	}

	// 4. Insert the new user into the database
	res, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// On success, return the new user data
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": newUser, "insertedID": res.InsertedID})
}
