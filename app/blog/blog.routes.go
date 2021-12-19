package blog

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/", ArticlesCategoryHandler).Methods("GET")
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", "testaaae")
}
