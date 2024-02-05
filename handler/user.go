package handler

import (
	"net/http"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type User struct {
	l hclog.Logger
}

func NewUser(l hclog.Logger) *User {
	return &User{l}
}

func (u *User) Authenticate(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	// Validate credentials
	if !isValidCredentials(user) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := generateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set token in the response header
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"jwt": token})
}
