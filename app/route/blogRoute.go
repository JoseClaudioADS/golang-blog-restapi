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
	blogHandler := handler.NewBlogHandler(&blogRepository)

	router.HandleFunc("", blogHandler.CreateHandle).Methods(http.MethodPost)
	router.HandleFunc("", blogHandler.ListHandle).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", blogHandler.GetHandle).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", blogHandler.DeleteHandle).Methods(http.MethodDelete)

}
