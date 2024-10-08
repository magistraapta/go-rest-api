package handlers

import (
	"example/hello/internal/db"
	"example/hello/internal/models"
	"example/hello/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func GetAllPosts(c echo.Context) error {
	var posts []models.Post
	err := db.DB.Select(&posts, "SELECT id, title, content, created_at FROM posts")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func GetPostById(c echo.Context) error {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	var post models.Post
	err = db.DB.Get(&post, "SELECT * FROM posts WHERE id = $1", postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) error {

	title := c.FormValue("title")
	content := c.FormValue("content")

	_, err := db.DB.Exec("INSERT INTO posts (title, content) VALUES ($1, $2)", title, content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func UpdatePost(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	var post models.Post
	err = c.Bind(&post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	post.UpdatedAt = time.Now()

	_, err = db.DB.NamedExec("UPDATE posts SET title = :title, content = :content, updated_at = :updated_at WHERE id = :id", post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, post)
}

func DeletePost(c echo.Context) error {
	id := c.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	_, err = db.DB.Exec("DELETE FROM posts WHERE id = $1", postId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response.SuccessResponse)
}

func RenderPostPage(c echo.Context) error {
	return c.Render(http.StatusOK, "create_post", nil)
}
