package blog

import (
	"net/http"

	blog "github.com/JoseClaudioADS/golang-blog-restapi/app/blog/handler"
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("", blog.CreateBlogHandler).Methods(http.MethodPost)
}
