package postmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID       primitive.ObjectID `json:"id"       bson:"_id,omitempty"`
	FileData []string           `json:"fileData" bson:"fileData"` // <-- FIXED
	Message  string             `json:"message"  bson:"message"`
	UserID   primitive.ObjectID `json:"userId"   bson:"userId"` // foreign key
}
