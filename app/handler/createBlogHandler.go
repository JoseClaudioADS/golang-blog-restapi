package handler

import (
	"encoding/json"
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app/domain"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/service"
	"github.com/go-playground/validator/v10"
)

type CreateBlogRequest struct {
	Title       string `validate:"required,gte=4,lte=200"`
	Description string `validate:"required,gte=2,lte=200"`
	Author      string `validate:"required,gte=2,lte=250"`
}

func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {

	var createBlogRequest CreateBlogRequest

	err := json.NewDecoder(r.Body).Decode(&createBlogRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(createBlogRequest)
	if validateErr != nil {
		// validationErrors := validateErr.(validator.ValidationErrors)
		// http.Error(w, validationErrors.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blog := domain.Blog{
		Title:       createBlogRequest.Title,
		Description: createBlogRequest.Description,
		Author:      createBlogRequest.Author,
	}

	service.CreateBlogService{}.Execute(blog)
	w.WriteHeader(http.StatusCreated)
}
