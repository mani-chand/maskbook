package user

import (
	"bytes"
	"context"
	"fmt"

	// "errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	// 1. THIS IS THE CORRECT IMPORT.
	// It must match the import in main.go
	"noob/services/database"
	"noob/services/storage"

	// 2. Use an alias for your models
	usermodels "noob/models/user"
	"noob/utils"
)

type SignupInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	FileData string `json:"fileData"`
}

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
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	existingUser.Password = ""
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": existingUser})
}

// CreateUser handles the logic for creating a new user.
func CreateUser(c *gin.Context) {
	// BUG FIX 1: Bind to AuthInput to get the plain-text password
	var input SignupInput
	fileURL := ""

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
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	if input.FileData != "" {
		mimeType, err := utils.GetMimeType(input.FileData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 2. Get the correct file extension (e.g., ".jpg")
		extension, err := utils.MimeTypeToExtension(mimeType)
		if err != nil {
			// This means the user uploaded an unsupported file type
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 3. Decode the data
		fileData, err := utils.DecodeBase64File(input.FileData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 4. Create the new, correct filename
		filename := uuid.New().String() + extension
		fileURL, err = storage.UploadFile(filename, bytes.NewReader(fileData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to upload file %s: %s", filename, err),
			})
			return
		}
		log.Println("File uploaded to S3:", fileURL)
	} else {
		log.Println("No file data was sent with this post.")
	}

	// 3. Create a new user object
	newUser := usermodels.User{
		ID:       primitive.NewObjectID(),
		Username: input.Username,
		Email:    input.Email,
		Password: string(hash), // from bcrypt
		Mobile:   input.Mobile,
		Avatar:   fileURL,
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
