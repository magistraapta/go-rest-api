package handlers

import (
	"example/hello/internal/db"
	"example/hello/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardPageData struct {
	Title string
	Users []models.User
}

func DashboardViewHandler(c echo.Context) error {
	var users []models.User

	err := db.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := DashboardPageData{
		Title: "Dashboard",
		Users: users,
	}
	return c.Render(http.StatusOK, "dashboard", data)
}
