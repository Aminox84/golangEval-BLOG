package models

type User struct {
	ID       int    `db:"id"`
	Lastname string `db:"lastname"`
	Email    string `db:"email"`
}
