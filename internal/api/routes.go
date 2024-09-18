package api

import (
	"example/hello/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/users", handlers.GetAllUsers)
	e.GET("/user/:id", handlers.GetUser)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/update/:id", handlers.UpdateUser)
	e.DELETE("/users/delete/:id", handlers.DeleteUser)

	e.GET("/books", handlers.GetAllBooks)
	e.POST("/books", handlers.CreateBook)
	e.GET("/book/:id", handlers.GetBookById)
	e.PUT("/book/update/:id", handlers.UpdateBook)
	e.DELETE("/book/delete/:id", handlers.DeleteBook)
}
