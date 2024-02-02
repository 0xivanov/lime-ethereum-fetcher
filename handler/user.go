package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/hashicorp/go-hclog"
)

type User struct {
	l hclog.Logger
	v *validator.Validate
}

func NewUser(l hclog.Logger, v *validator.Validate) *User {
	return &User{l, v}
}

func (u *User) Authenticate(c *gin.Context) {
	// TODO
}
