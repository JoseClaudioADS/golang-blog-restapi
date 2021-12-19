package route

import (
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app/handler"
	"github.com/gorilla/mux"
)

func InitBlogRoutes(router *mux.Router) {
	router.HandleFunc("", handler.CreateBlogHandler).Methods(http.MethodPost)
}
