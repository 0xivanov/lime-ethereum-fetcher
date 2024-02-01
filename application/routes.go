package application

import (
	"github.com/gin-gonic/gin"
)

func (app *App) loadRoutes() {
	app.r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
