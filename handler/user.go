package handler

import (
	"net/http"
	"time"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
)

type User struct {
	l         hclog.Logger
	tr        repo.TransactionInterface
	JwtSecret []byte
}

func NewUser(l hclog.Logger, tr repo.TransactionInterface, jwtSecret string) *User {
	return &User{l, tr, []byte(jwtSecret)}
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
	token, err := u.generateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set token in the response header
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// isValidCredentials checks if the given credentials are valid
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

func (u *User) generateToken(username string) (string, error) {
	claims := &JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(u.JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uh *User) GetUserTransactions(c *gin.Context) {
	authToken := c.GetHeader("AUTH_TOKEN")
	if authToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing AUTH_TOKEN header"})
		return
	}
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return uh.JwtSecret, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			return
		}

		// Extract the username from the claims
		username := claims["username"].(string)
		transactions, err := uh.tr.GetTransactionsByUsername(c, username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load transactions for this username: " + username})
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
		return
	}

}
