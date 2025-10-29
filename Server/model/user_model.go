package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID          string        `bson:"user_id" json:"user_id"`
	FirstName       string        `bson:"first_name" json:"first_name" validation:"required,min=2,max=100"`
	LastName        string        `bson:"last_name" json:"last_name" validation:"required,min=2,max=100"`
	Email           string        `bson:"email" json:"email" validation:"required,email"`
	Password        string        `bson:"password" json:"password" validation:"required,min=6"`
	Role            string        `bson:"role" json:"role" validation:"oneof=ADMIN USER"`
	CreatedAt       time.Time     `bson:"created_at" json:"created_at" validation:"required"`
	UpdatedAt       time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	Token           string        `bson:"token" json:"token"`
	RefreshToken    string        `bson:"refresh_token" json:"refresh_token"`
	FavouriteGenres []Genre       `bson:"favourite_genres" json:"favourite_genres" validation:"required,dive"`
}

type UserLogin struct {
	Email    string `bson:"email" json:"email" validation:"required,email"`
	Password string `bson:"password" json:"password" validation:"required,min=6"`
}

type UserResponse struct {
	UserID          string  `json:"user_id"`
	FirstName       string  `json:"first_name" validation:"required,min=2,max=100"`
	LastName        string  `json:"last_name" validation:"required,min=2,max=100"`
	Email           string  `json:"email" validation:"required,email"`
	Role            string  `json:"role" validation:"oneof=ADMIN USER"`
	Token           string  `json:"token"`
	RefreshToken    string  `json:"refresh_token"`
	FavouriteGenres []Genre `json:"favourite_genres" validation:"required,dive"`
}
