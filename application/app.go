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
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/hashicorp/go-hclog"
)

type App struct {
	r *gin.Engine
	p string
	l hclog.Logger
}

func New(r *gin.Engine, p string, db *db.Database, l hclog.Logger, v *validator.Validate, tr *repo.Transaction) *App {
	app := &App{r, p, l}
	app.loadRoutes(handler.NewTransaction(l, v, tr), handler.NewUser(l, v))
	return app
}

func (app *App) Start() {
	s := http.Server{
		Addr:         "localhost:" + app.p, // configure the bind address
		Handler:      app.r,                // set the default handler
		ReadTimeout:  5 * time.Second,      // max time to read request from the client
		WriteTimeout: 10 * time.Second,     // max time to write response to the client
		IdleTimeout:  120 * time.Second,    // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		app.l.Info("Starting server on", "port", app.p)

		err := s.ListenAndServe()
		if err != nil {
			app.l.Error("Starting server", "error", err)
			os.Exit(1)
		}
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
