package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Username  string               `json:"username" bson:"username" binding:"required"`
	Email     string               `json:"email" bson:"email" binding:"required,email"`
	Password  string               `json:"password,omitempty" bson:"password" binding:"required"`
	Name      string               `json:"name" bson:"name"`
	Bio       string               `json:"bio" bson:"bio"`
	Avatar    string               `json:"avatar" bson:"avatar"`
	Following []primitive.ObjectID `json:"following" bson:"following"`
	Followers []primitive.ObjectID `json:"followers" bson:"followers"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time            `json:"updated_at" bson:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
