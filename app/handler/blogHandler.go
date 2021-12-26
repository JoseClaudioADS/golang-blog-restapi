package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/JoseClaudioADS/golang-blog-restapi/app/model"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type blogHandler struct {
	blogRepository *repository.BlogRepository
}

func NewBlogHandler(blogRepository *repository.BlogRepository) blogHandler {
	return blogHandler{
		blogRepository: blogRepository,
	}
}

func (c blogHandler) CreateHandle(w http.ResponseWriter, r *http.Request) {

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

	blog := model.Blog{
		Title:       createBlogRequest.Title,
		Description: createBlogRequest.Description,
		Author:      createBlogRequest.Author,
	}
	(*c.blogRepository).Create(&blog)
	w.WriteHeader(http.StatusCreated)
}

func (c blogHandler) GetHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	blog, err := (*c.blogRepository).GetById(&id)
	if err != nil {
		if blog.Id == 0 {
			SendResponse(w, nil, http.StatusNotFound)
		} else {
			log.Println("error on get blog", err)
			SendResponse(w, nil, http.StatusInternalServerError)
		}
		return
	}

	responseData := GetBlogResponseItem{
		Id:          blog.Id,
		Title:       blog.Title,
		Description: blog.Description,
		Author:      blog.Author,
	}

	SendResponse(w, responseData, http.StatusOK)
}

func (c blogHandler) ListHandle(w http.ResponseWriter, r *http.Request) {

	blogs, err := (*c.blogRepository).List()
	if err != nil {
		log.Println("error on list blogs", err)
	}

	var responseData []ListBlogResponseItem

	for i := range blogs {
		blog := blogs[i]

		responseData = append(responseData, ListBlogResponseItem{
			Id:          blog.Id,
			Title:       blog.Title,
			Description: blog.Description,
			Author:      blog.Author,
		})
	}

	SendResponse(w, responseData, http.StatusOK)
}

func (c blogHandler) DeleteHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := (*c.blogRepository).DeleteById(&id)
	if err != nil {
		log.Println("error on delete blog", err)
	}

	SendResponse(w, nil, http.StatusOK)
}
