package service

import (
	"fmt"

	"github.com/JoseClaudioADS/golang-blog-restapi/app/domain"
)

type CreateBlogService struct {
}

func (c CreateBlogService) Execute(blog domain.Blog) (domain.Blog, error) {
	fmt.Println(blog)
	return blog, nil
}
