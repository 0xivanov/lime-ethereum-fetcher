package application

import (
	"github.com/0xivanov/lime-ethereum-fetcher-go/handler"
	"github.com/gin-gonic/gin"
)

func (app *App) loadRoutes(th *handler.Transaction, uh *handler.User) {

	app.r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.r.GET("/lime/all", th.GetTransactions)
	app.r.GET("/lime/eth", th.GetTransactionsWithHashes)
	app.r.POST("/lime/authenticate", uh.Authenticate)
}
