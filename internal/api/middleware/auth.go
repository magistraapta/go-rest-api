package middleware

import (
	"example/hello/internal/api/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("Token")
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		claims, err := handlers.VerifyToken(cookie.Value)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		c.Set("user", claims.Username)
		return next(c)
	}
}
