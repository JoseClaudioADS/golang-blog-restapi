package app

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	initRoutes(app)
	return app
}
