package app

import "github.com/JoseClaudioADS/golang-blog-restapi/app/route"

func initRoutes(app *App) {
	blogRouter := app.Router.PathPrefix("/blog").Subrouter()
	route.InitBlogRoutes(blogRouter)
}
