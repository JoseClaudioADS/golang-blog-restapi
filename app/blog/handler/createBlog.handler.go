package blog

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CreateBlogInput struct {
	title       string `validate:"required,gte=4,lte=200"`
	description string `validate:"required,gte=2,lte=200"`
	author      string `validate:"required,gte=2,lte=250"`
}

func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {

	var createBlogInput CreateBlogInput

	err := json.NewDecoder(r.Body).Decode(&createBlogInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(createBlogInput)
	if validateErr != nil {
		// validationErrors := validateErr.(validator.ValidationErrors)
		// http.Error(w, validationErrors.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
