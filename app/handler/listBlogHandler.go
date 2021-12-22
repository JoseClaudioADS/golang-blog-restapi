package handler

import (
	"log"
	"net/http"

	"github.com/JoseClaudioADS/golang-blog-restapi/app/repository"
)

type listBlogResponseItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type listBlogHandler struct {
	blogRepository *repository.BlogRepository
}

func NewListBlogHandler(blogRepository *repository.BlogRepository) listBlogHandler {
	return listBlogHandler{
		blogRepository: blogRepository,
	}
}

func (c listBlogHandler) Handle(w http.ResponseWriter, r *http.Request) {

	blogs, err := (*c.blogRepository).List()
	if err != nil {
		log.Println("error on list blogs", err)
	}

	var blogItems []listBlogResponseItem

	for i := range blogs {
		blog := blogs[i]

		blogItems = append(blogItems, listBlogResponseItem{
			Id:          blog.Id,
			Title:       blog.Title,
			Description: blog.Description,
			Author:      blog.Author,
		})
	}

	responseData := make(map[string][]listBlogResponseItem)
	responseData["blogs"] = blogItems

	SendResponse(w, responseData, http.StatusOK)
}
