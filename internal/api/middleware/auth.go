package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(username, password string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")

			if auth == "" {
				return c.String(http.StatusUnauthorized, "Unautorized")
			}

			if !strings.HasPrefix(auth, "Basic") {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic"))

			if err != nil {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			parts := strings.SplitN(string(payload), ":", 2)
			if len(parts) != 2 || parts[0] != username || parts[1] != password {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	}
}
