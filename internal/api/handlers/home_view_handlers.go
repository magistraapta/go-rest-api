package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"example/hello/internal/db"
	"example/hello/internal/models"
)

type PageData struct {
	Title string
	Posts []models.Post
}

func HomeHandler(c echo.Context) error {
	var posts []models.Post
	err := db.DB.Select(&posts, "SELECT id, title, content, created_at FROM posts ORDER BY created_at DESC LIMIT 10")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching posts")
	}

	data := PageData{
		Title: "Home",
		Posts: posts,
	}

	return c.Render(http.StatusOK, "home", data)
}
