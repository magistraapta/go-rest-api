package handlers

import (
	"example/hello/internal/db"
	"example/hello/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DetailPageData struct {
	Title string
	Post  models.Post
}

func DetailHandler(c echo.Context) error {
	// Get the post ID from the URL parameters
	postID := c.Param("id")

	var post models.Post
	err := db.DB.Get(&post, "SELECT * FROM posts WHERE id = $1", postID)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching post")
	}

	data := DetailPageData{
		Title: post.Title,
		Post:  post,
	}

	return c.Render(http.StatusOK, "detail_post", data)
}
