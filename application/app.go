package application

import (
	"log"

	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/gin-gonic/gin"
)

type App struct {
	r    *gin.Engine
	port string
	l    *log.Logger
	// v *validator.Validate
}

func New(r *gin.Engine, port string, db *db.Database, l *log.Logger) *App {
	app := &App{r, port, l}
	app.loadRoutes()
	return app
}

func (app *App) Start() {
	app.r.Run("localhost:" + app.port)
}
