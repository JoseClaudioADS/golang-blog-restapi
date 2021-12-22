package model

type Blog struct {
	Title       string `db:"title"`
	Description string `db:"description"`
	Author      string `db:"author"`
}
