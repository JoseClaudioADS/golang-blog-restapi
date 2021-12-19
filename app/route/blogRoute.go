package route

import (
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app/handler"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/infra/database"
	"github.com/gorilla/mux"
)

func InitBlogRoutes(router *mux.Router, db database.DB) {
	router.HandleFunc("", handler.CreateBlogHandler{DB: db}.Handle).Methods(http.MethodPost)
}
