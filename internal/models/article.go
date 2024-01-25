package models

type Article struct {
	ID       int    `db:"id"`
	Content  string `db:"content"`
	Likes    int    `db:"likes"`
	Dislikes int    `db:"dislikes"`
}
