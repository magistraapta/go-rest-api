package models

import (
	"html/template"
	"time"
)

type Post struct {
	ID        int           `db:"id" json:"id"`
	Title     string        `db:"title" json:"title"`
	Content   template.HTML `db:"content" json:"content"`
	CreatedAt time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" json:"updated_at"`
}
