package user

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

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

// decodeBase64File is a helper function to strip the prefix
// (e.g., "data:image/png;base64,") and decode the data.
func decodeBase64File(dataURL string) ([]byte, error) {
	// Find the comma
	commaIndex := strings.Index(dataURL, ",")
	if commaIndex == -1 {
		return nil, errors.New("invalid base64 data URL: missing comma")
	}

	// Get the part after the comma
	base64Data := dataURL[commaIndex+1:]

	// Decode the string
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, errors.New("failed to decode base64 data")
	}
	return decoded, nil
}

// --- HELPER FUNCTIONS ---

// getMimeType parses the data URL (e.g., "data:image/png;base64,...")
// and returns the MIME type (e.g., "image/png")
func getMimeType(dataURL string) (string, error) {
	// Find "data:"
	startIndex := strings.Index(dataURL, "data:")
	if startIndex == -1 {
		return "", errors.New("invalid data URL: missing 'data:' prefix")
	}
	// Find ";base64,"
	endIndex := strings.Index(dataURL, ";base64,")
	if endIndex == -1 {
		return "", errors.New("invalid data URL: missing ';base64,' separator")
	}

	// The MIME type is between "data:" and ";base64,"
	mimeType := dataURL[startIndex+5 : endIndex]
	return mimeType, nil
}

// mimeTypeToExtension maps a MIME type to a file extension.
// Add any other file types you want to support here.
func mimeTypeToExtension(mimeType string) (string, error) {
	switch mimeType {
	case "image/jpeg":
		return ".jpg", nil
	case "image/png":
		return ".png", nil
	case "image/gif":
		return ".gif", nil
	case "video/mp4":
		return ".mp4", nil
	case "video/quicktime":
		return ".mov", nil
	default:
		return "", errors.New("unsupported file type: " + mimeType)
	}
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
		mimeType, err := getMimeType(input.FileData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 2. Get the correct file extension (e.g., ".jpg")
		extension, err := mimeTypeToExtension(mimeType)
		if err != nil {
			// This means the user uploaded an unsupported file type
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 3. Decode the data
		fileData, err := decodeBase64File(input.FileData)
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
