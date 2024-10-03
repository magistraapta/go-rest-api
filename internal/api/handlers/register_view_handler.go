package handlers

import (
	"example/hello/internal/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterViewHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "register", nil)
}

func RegisterHandler(c echo.Context) error {
	username := c.FormValue("username")
	passowrd := c.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passowrd), bcrypt.DefaultCost)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error hashing password")
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error inserting user into database")
	}

	return c.String(http.StatusCreated, "user registered successfully")
}
