package handlers

import (
	"example/hello/internal/db"
	"example/hello/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardPageData struct {
	Title string
	Posts []models.Post
	Users []models.User
}

func DashboardViewHandler(c echo.Context) error {
	var posts []models.Post
	var users []models.User

	err_post := db.DB.Select(&posts, "SELECT * FROM posts")
	if err_post != nil {
		return c.String(http.StatusInternalServerError, err_post.Error())
	}

	err_user := db.DB.Select(&users, "SELECT * FROM users")
	if err_user != nil {
		return c.String(http.StatusInternalServerError, err_user.Error())
	}


	data := DashboardPageData{
		Title: "Dashboard",
		Posts: posts,
		Users: users,
	}
	return c.Render(http.StatusOK, "dashboard", data)
}
