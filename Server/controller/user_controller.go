package controller

import (
	"context"
	"mMoviez/database"
	"mMoviez/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = database.OpenCollection("users")

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		err := checker.Struct(user)

		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while checking for existing user"})
			return
		}
		if count > 0 {
			c.JSON(400, gin.H{"error": "User with this email already exists"})
			return
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(500, gin.H{"error": "Error while hashing password"})
			return
		}

		user.Password = string(hashed)
		user.CreatedAt = time.Now()
		user.UpdatedAt = user.CreatedAt

		result, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while inserting user"})
			return
		}
		c.JSON(201, result)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cursor, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while fetching users"})
			return
		}

		var users []model.User
		if err = cursor.All(ctx, &users); err != nil {
			c.JSON(500, gin.H{"error": "Error while parsing users"})
			return
		}
		c.JSON(200, users)
	}
}

func GetUserByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		email := c.Param("email")

		if email == "" {
			c.JSON(400, gin.H{"error": "Email is required"})
			return
		}

		var user model.User

		err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		id := c.Param("user_id")
		if id == "" {
			c.JSON(400, gin.H{"error": "User ID is required"})
			return
		}

		var user model.User

		err := userCollection.FindOne(ctx, bson.M{"user_id": id}).Decode(&user)

		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	}
}
