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
	"github.com/hashicorp/go-hclog"
)

type App struct {
	r *gin.Engine
	p string
	l hclog.Logger
	s *http.Server
}

func New(r *gin.Engine, p string, db *db.Database, l hclog.Logger, tr repo.TransactionInterface) *App {
	app := &App{r, p, l, nil}
	app.loadRoutes(handler.NewTransaction(l, tr), handler.NewUser(l))
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
	app.s = &s
	// start the server
	go func() {
		app.l.Info("Starting server on", "port", app.p)

		err := s.ListenAndServe()
		if err != nil {
			app.l.Info("Closing server", "error", err)
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

func (app *App) Stop() {
	if app.r != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := app.s.Shutdown(ctx); err != nil {
			app.l.Error("Shutdown server", "error", err)
		}
		app.l.Info("Server stopped gracefully")
	}
}
