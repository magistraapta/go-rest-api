package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PageData struct {
	Title   string
	Content string
}

func HomeHandler(c echo.Context) error {
	data := PageData{
		Title:   "Home",
		Content: "Welcome to the home page",
	}

	return c.Render(http.StatusOK, "home.html", data)
}
