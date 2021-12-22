package app

import (
	"github.com/JoseClaudioADS/golang-blog-restapi/app/infra/database"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *database.DB
}

func New(db *database.DB) *App {
	app := &App{
		Router: mux.NewRouter(),
		DB:     db,
	}

	initRoutes(app)
	return app
}
