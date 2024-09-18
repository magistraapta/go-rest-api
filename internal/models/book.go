package models

type Book struct {
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}
