package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "auth", nil)
}

func HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "admin" && password == "password" {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.String(http.StatusUnauthorized, "Invalid username or password")
}
