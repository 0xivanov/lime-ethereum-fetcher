package application

import (
	"github.com/0xivanov/lime-ethereum-fetcher-go/handler"
	"github.com/gin-gonic/gin"
)

func (app *App) loadRoutes(th *handler.Transaction, uh *handler.User) {

	// basic endpoint to test if the app is running properly
	app.r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// assignment lime endpoints
	app.r.GET("/lime/all", th.GetTransactions)
	app.r.GET("/lime/eth", handler.AuthenticateMiddleware(), th.GetTransactionsWithHashes)
	app.r.GET("/lime/eth/:rlphex", handler.AuthenticateMiddleware(), th.GetTransactionsWithRlp)
	app.r.POST("/lime/authenticate", uh.Authenticate)
}
