package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Username  string             `json:"username" bson:"username"`
	Content   string             `json:"content" bson:"content" binding:"required,max=280"`
	Likes     []primitive.ObjectID `json:"likes" bson:"likes"`
	Retweets  []primitive.ObjectID `json:"retweets" bson:"retweets"`
	Replies   []primitive.ObjectID `json:"replies" bson:"replies"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateTweetRequest struct {
	Content string `json:"content" binding:"required,max=280"`
}
