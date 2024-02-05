package handler

import (
	"net/http"
	"time"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("test-key")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			c.Abort()
			return
		}

		// Validate credentials
		if !isValidCredentials(user) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			c.Abort()
			return
		}

		// Generate JWT token
		token, err := generateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			c.Abort()
			return
		}

		// Set token in the response header
		c.Header("Authorization", "Bearer "+token)
		c.Next()
	}
}

func isValidCredentials(user model.User) bool {
	validCredentials := map[string]string{
		"alice": "alice",
		"bob":   "bob",
		"carol": "carol",
		"dave":  "dave",
	}
	password, ok := validCredentials[user.Username]
	if !ok || password != user.Password {
		return false
	}
	return true
}

func generateToken(username string) (string, error) {
	claims := &JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}