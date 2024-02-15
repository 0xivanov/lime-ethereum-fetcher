package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/0xivanov/lime-ethereum-fetcher-go/handler"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type App struct {
	router        *gin.Engine
	port          string
	logger        hclog.Logger
	server        *http.Server
	client        *ethclient.Client
	wsClient      *ethclient.Client
	eventListener *EventListener
	ctx           context.Context
	cancelFunc    context.CancelFunc
}

func New(r *gin.Engine, p string, ethNodeUrl string, jwtSecret string, client *ethclient.Client, wsClient *ethclient.Client, db *db.Database, l hclog.Logger, tr repo.TransactionInterface, cr repo.ContractInterface) *App {
	ctx, cancel := context.WithCancel(context.Background())
	app := &App{r, p, l, nil, client, wsClient, &EventListener{l, cr}, ctx, cancel}
	app.loadRoutes(handler.NewTransaction(l, tr, ethNodeUrl), handler.NewUser(l, tr, jwtSecret), handler.NewSmartContract(l, client, cr))
	return app
}

func (app *App) Start() {
	s := http.Server{
		Addr:         "0.0.0.0:" + app.port, // configure the bind address
		Handler:      app.router,            // set the default handler
		ReadTimeout:  5 * time.Second,       // max time to read request from the client
		WriteTimeout: 10 * time.Second,      // max time to write response to the client
		IdleTimeout:  120 * time.Second,     // max time for connections using TCP Keep-Alive
	}
	app.server = &s
	// start the server
	go func() {
		app.logger.Info("Starting server on", "port", app.port)

		err := s.ListenAndServe()
		if err != nil {
			app.logger.Info("Closing server", "error", err)
		}
	}()

	// run event listener
	go func() {
		defer func() {
			if r := recover(); r != nil {
				app.logger.Error("event listener goroutine panicked", "error", r)
			}
		}()
		app.eventListener.PersonInfoEventListenerStart(app.ctx, app.wsClient)
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

func (app *App) Stop() {
	// Cancel the context to stop all goroutines
	app.cancelFunc()

	if app.router != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := app.server.Shutdown(ctx); err != nil {
			app.logger.Error("Shutdown server", "info", err)
		}
		app.logger.Info("Server stopped gracefully")
	}
}
