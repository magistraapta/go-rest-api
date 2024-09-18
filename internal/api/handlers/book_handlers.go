package handlers

import (
	"net/http"
	"strconv"

	"example/hello/internal/db"
	"example/hello/internal/models"
	"example/hello/pkg/response"

	"github.com/labstack/echo/v4"
)

func GetAllBooks(c echo.Context) error {
	var books []models.Book
	err := db.DB.Select(&books, "SELECT * FROM book")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, books)
}

func GetBookById(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	book := models.Book{}
	err = db.DB.Get(&book, "SELECT * FROM book WHERE id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, book)
}

func CreateBook(c echo.Context) error {
	book := models.Book{}
	err := c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	_, err = db.DB.NamedExec("INSERT INTO book (title, author) VALUES (:title, :author)", book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	book := models.Book{}
	err = c.Bind(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	_, err = db.DB.NamedExec("UPDATE book SET title = :title, author = :author WHERE id = :id", book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}

	return c.JSON(http.StatusOK, response.SuccessResponse)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse)
	}

	_, err = db.DB.Exec("DELETE FROM book WHERE id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse)
}

// Implement other book handlers (CreateBook, GetSpecificBook, UpdateBook, DeleteBook) similarly
