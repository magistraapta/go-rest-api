package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DashboardViewHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard", nil)
}
