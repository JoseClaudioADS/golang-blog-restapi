package repository

import (
	"github.com/JoseClaudioADS/golang-blog-restapi/app/infra/database"
	"github.com/JoseClaudioADS/golang-blog-restapi/app/model"
)

type BlogRepository interface {
	Create(blog *model.Blog) error
	List() ([]model.Blog, error)
}

type blogRepositoryDatabase struct {
	db *database.DB
}

func NewBlogRepositoryDatabase(db *database.DB) BlogRepository {
	return blogRepositoryDatabase{
		db: db,
	}
}

func (b blogRepositoryDatabase) Create(blog *model.Blog) error {
	_, err := b.db.Connection.NamedExec("INSERT INTO BLOG (title, description, author) VALUES (:title, :description, :author)", &blog)
	return err
}

func (b blogRepositoryDatabase) List() ([]model.Blog, error) {
	blogs := []model.Blog{}
	err := b.db.Connection.Select(&blogs, "SELECT * FROM BLOG")
	return blogs, err
}
