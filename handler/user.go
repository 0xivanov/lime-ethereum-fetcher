package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type User struct {
	l *log.Logger
	v *validator.Validate
}

func NewUser(l *log.Logger, v *validator.Validate) *User {
	return &User{l, v}
}

func (u *User) CreateUser(c *gin.Context) {
	// var newUser model.User

	// // Bind the request body to the newUser variable
	// if err := c.ShouldBindJSON(&newUser); err != nil {
	// 	u.l.Printf("Error binding JSON: %s", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Validate the user struct
	// if err := u.v.Struct(newUser); err != nil {
	// 	u.l.Printf("Validation error: %s", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// newUser.ID = len(model.GetUsers()) + 1
	// newUser.Persist()

	// u.l.Printf("User created successfully. ID: %d, Username: %s", newUser.ID, newUser.Username)

	// // Respond with the created user
	// c.JSON(http.StatusCreated, newUser)
}

func (u *User) GetUsers(c *gin.Context) {
	// users := model.GetUsers()

	// u.l.Printf("Retrieved %d users", len(users))

	// c.JSON(http.StatusOK, users)
}
