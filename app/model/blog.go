package model

type Blog struct {
	Id          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Author      string `db:"author"`
}
