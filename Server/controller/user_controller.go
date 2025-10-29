package controller

import (
	"context"
	"mMoviez/database"
	"mMoviez/model"
	"mMoviez/util"
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

		user.UserID = bson.NewObjectID().Hex()
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

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userLogin model.UserLogin

		if err := c.ShouldBindJSON(&userLogin); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		err := checker.Struct(userLogin)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var foundUser model.User
		err = userCollection.FindOne(ctx, bson.M{"email": userLogin.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(userLogin.Password))
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		token, refreshToken, err := util.GenerateAllTokens(foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserID, foundUser.Role)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while generating tokens"})
			return
		}

		err = util.UpdateAllTokens(foundUser.UserID, token, refreshToken)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while updating tokens"})
			return
		}

		c.JSON(200,
			&model.UserResponse{
				UserID:          foundUser.UserID,
				FirstName:       foundUser.FirstName,
				LastName:        foundUser.LastName,
				Email:           foundUser.Email,
				Role:            foundUser.Role,
				Token:           token,
				RefreshToken:    refreshToken,
				FavouriteGenres: foundUser.FavouriteGenres,
			},
		)
	}
}
