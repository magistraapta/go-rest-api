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
	err := db.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&user.Username, &user.Password)

	if err != nil {
		// Any other database error
		return c.String(http.StatusInternalServerError, "Error querying the database")
	}

	// Compare the provided password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// If the password doesn't match
		return c.String(http.StatusUnauthorized, "Invalid username or password")
	}

	// If the login is successful, redirect to home page
	return c.Redirect(http.StatusSeeOther, "/")
}
