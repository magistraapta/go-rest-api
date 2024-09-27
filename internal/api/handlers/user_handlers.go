package handlers

import (
	"net/http"
	"strconv"

	"example/hello/internal/db"
	"example/hello/internal/models"
	"example/hello/pkg/response"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	var users []models.User
	err := db.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	var user models.User
	err = db.DB.Get(&user, "SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	_, err := db.DB.NamedExec("INSERT INTO users (name, phone, address) VALUES (:name, :phone, :address)", user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}

	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	_, err = db.DB.NamedExec("UPDATE users SET name = :name, phone = :phone, address = :address WHERE id = :id", user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	_, err = db.DB.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse)
}

