package handlers

import (
	"example/hello/internal/db"
	"example/hello/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	errDb := db.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&user.Username, &user.Password)

	tokenString, expirationTime, errToken := CreateToken(user.Username)
	if errDb != nil {
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	if errToken != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate token")
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errHash != nil {
		return c.String(http.StatusUnauthorized, "Invalid username or password")
	}

	c.SetCookie(&http.Cookie{
		Name:    "Token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return c.Redirect(http.StatusSeeOther, "/")
}
