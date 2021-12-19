package app

import (
	"github.com/JoseClaudioADS/golang-blog-restapi/app/infra"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	db := infra.DB{}
	db.Open()
	defer db.Close()

	initRoutes(app)
	return app
}
