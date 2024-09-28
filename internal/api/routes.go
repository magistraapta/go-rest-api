package api

import (
	"example/hello/internal/api/handlers"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func SetupRoutes(e *echo.Echo) {

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("../../views/*.html")),
	}
	e.Renderer = renderer

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

	e.GET("/posts", handlers.GetAllPosts)
	e.POST("/posts", handlers.CreatePost)
	e.GET("/post/:id", handlers.GetPostById)
	e.PUT("/post/update/:id", handlers.UpdatePost)
	e.DELETE("/post/delete/:id", handlers.DeletePost)

	e.GET("/", handlers.HomeHandler)
	e.GET("/detail/:id", handlers.DetailHandler)
}
