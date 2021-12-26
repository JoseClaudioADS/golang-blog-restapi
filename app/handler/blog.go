package handler

type CreateBlogRequest struct {
	Title       string `validate:"required,gte=4,lte=200"`
	Description string `validate:"required,gte=2,lte=200"`
	Author      string `validate:"required,gte=2,lte=250"`
}

type ListBlogResponseItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type GetBlogResponseItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
