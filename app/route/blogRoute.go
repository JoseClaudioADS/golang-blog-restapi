package route

import (
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app/handler"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/infra/database"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/repository"
	"github.com/gorilla/mux"
)

func InitBlogRoutes(router *mux.Router, db *database.DB) {
	blogRepository := repository.NewBlogRepositoryDatabase(db)

	newBlogHandler := handler.NewBlogHandler(&blogRepository)
	router.HandleFunc("", newBlogHandler.Handle).Methods(http.MethodPost)

	newListBlogHandler := handler.NewListBlogHandler(&blogRepository)
	router.HandleFunc("", newListBlogHandler.Handle).Methods(http.MethodGet)

}
