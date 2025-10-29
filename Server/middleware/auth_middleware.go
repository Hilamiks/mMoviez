package middleware

import (
	"errors"
	"mMoviez/util"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) (string, error) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("Auth header required")
	}
	tokenString = tokenString[len("Bearer "):]
	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetAccessToken(c)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if token == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := util.ValidateToken(token)

		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
