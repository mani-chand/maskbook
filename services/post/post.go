package post

import (
	"bytes"
	"context"
	"net/http"
	"noob/services/database"
	"noob/services/storage"
	"noob/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // Import this
	"go.mongodb.org/mongo-driver/mongo"
)

type PostForm struct {
	FileData []string `json:"fileData"`
	UserId   string   `json:"user_id" bson:"user_id"`
	Message  string   `json:"message" bson:"message"`
	CreatedAt time.Time `json:"created_at" bson:"creation_time"`
}

// Assuming your struct looks like this based on previous context
// If you want the Mongo _id returned to frontend, use primitive.ObjectID
type Post struct {
	ID        interface{} `bson:"_id" json:"id"` // Use interface{} or primitive.ObjectID to handle the BSON ID
	Message   string      `bson:"message" json:"message"`
	FileData  []string    `bson:"fileData" json:"fileData"`
	UserId    string      `bson:"user_id" json:"user_id"`
	CreatedAt time.Time   `json:"created_at" bson:"creation_time"`
}
type AuthorDetails struct {
	Username string `bson:"username" json:"username"`
	Avatar   string `bson:"avatar"   json:"avatar"`
}

type PostWithUser struct {
	ID        interface{}   `bson:"_id"            json:"id"`
	Message   string        `bson:"message"        json:"message"`
	FileData  []string      `bson:"filedata"       json:"fileData"`
	UserId    string        `bson:"user_id"        json:"user_id"`
	Author    AuthorDetails `bson:"author_details" json:"author"` // This field comes from the lookup
	CreatedAt time.Time     `bson:"creation_time"  json:"created_at"`
}

func Get_all_posts(c *gin.Context) {
	db := database.GetDatabase()
	postCollection := db.Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the Aggregation Pipeline
	pipeline := mongo.Pipeline{
		// Stage 1: Sort by newest first
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: -1}}}},

		// Stage 2: Convert 'user_id' (String) to ObjectId so we can match it
		// We create a temporary field called 'userObjId'
		{{Key: "$addFields", Value: bson.D{
			{Key: "userObjId", Value: bson.D{{Key: "$toObjectId", Value: "$user_id"}}},
		}}},

		// Stage 3: Perform the Join (Lookup)
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},           // Target collection
			{Key: "localField", Value: "userObjId"}, // Field in 'posts' (our converted one)
			{Key: "foreignField", Value: "_id"},     // Field in 'users'
			{Key: "as", Value: "author_details"},    // Result field name
		}}},

		// Stage 4: Unwind the "author_details" array
		// (Lookup returns an array, this converts it to a single object)
		{{Key: "$unwind", Value: "$author_details"}},

		// Stage 5: Project (Optional - cleaning up fields we don't need)
		// Here we just ensure we drop the temporary 'userObjId' field
		{{Key: "$project", Value: bson.D{
			{Key: "userObjId", Value: 0},
		}}},
	}

	// Execute the Aggregation
	cursor, err := postCollection.Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}
	defer cursor.Close(ctx)

	// Decode results
	posts := []PostWithUser{}
	if err = cursor.All(ctx, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
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
	objID, err := primitive.ObjectIDFromHex(newPost.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID format"})
		return
	}

	// Validate user
	count, err := userCollection.CountDocuments(ctx, bson.M{"_id": objID})
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
	newPost.FileData = uploadedFiles
	newPost.CreatedAt = time.Now()

	postCollection := db.Collection("posts")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := postCollection.InsertOne(ctx, newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "post": newPost, "insertedID": res.InsertedID})

}
