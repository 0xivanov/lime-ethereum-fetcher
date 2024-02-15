package application

import (
	"github.com/0xivanov/lime-ethereum-fetcher-go/handler"
	"github.com/gin-gonic/gin"
)

func (app *App) loadRoutes(transactionHandler *handler.Transaction, userHandler *handler.User, smartContractHandler *handler.SmartContract) {

	// basic endpoint to test if the app is running properly
	app.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// lime endpoints
	app.router.GET("/lime/all", transactionHandler.GetTransactions)
	app.router.GET("/lime/eth", handler.AuthenticateMiddleware(), transactionHandler.GetTransactions)
	app.router.GET("/lime/eth/:rlphex", handler.AuthenticateMiddleware(), transactionHandler.GetTransactionsWithRlp)
	app.router.POST("/lime/authenticate", userHandler.Authenticate)
	app.router.GET("/lime/my", userHandler.GetUserTransactions)
	app.router.POST("/lime/savePerson", smartContractHandler.SavePerson)
	app.router.GET("/lime/listPersons", smartContractHandler.GetPersons)
}
