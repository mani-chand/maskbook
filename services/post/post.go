package post

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"noob/services/database"
	"noob/services/storage"
	"noob/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type PostForm struct {
	FileData []string `json:"fileData"`
	UserId   string   `json:"user_id" bson:"user_id"`
	Message  string   `json:"message" bson:"message"`
}

func Get_all_posts(c *gin.Context) {

}

func Create_post(c *gin.Context) {
	var newPost PostForm

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	db := database.GetDatabase()
	userCollection := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Validate user
	count, err := userCollection.CountDocuments(ctx, bson.M{"_id": newPost.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	// Stores URL of uploaded files
	uploadedFiles := []string{}

	// Process files
	if newPost.FileData != nil {
		for i := 0; i < len(newPost.FileData); i++ {
			file := newPost.FileData[i] // ✔️ FIXED

			mimeType, err := utils.GetMimeType(file)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			extension, err := utils.MimeTypeToExtension(mimeType)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			data, err := utils.DecodeBase64File(file)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			filename := uuid.New().String() + extension

			url, err := storage.UploadFile(filename, bytes.NewReader(data))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			uploadedFiles = append(uploadedFiles, url) // ✔️ SAVE URLs properly
		}
	}

	log.Println("Uploaded:", uploadedFiles)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post created",
		"files":   uploadedFiles,
	})
}
