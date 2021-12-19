package app

import "github.com/JoseClaudioADS/golang-blog-restapi/app/blog"

func initRoutes(app *App) {
	blogRouter := app.Router.PathPrefix("/blog").Subrouter()
	blog.InitRoutes(blogRouter)
}
