package handler

import (
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
	// TODO
}
