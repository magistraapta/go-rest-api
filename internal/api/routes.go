package api

import (
	"example/hello/internal/api/handlers"
	"example/hello/internal/api/middleware"
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

	e.GET("/posts", handlers.GetAllPosts)
	e.POST("/posts", handlers.CreatePost)
	e.GET("/post/:id", handlers.GetPostById)
	e.PUT("/post/update/:id", handlers.UpdatePost)
	e.DELETE("/post/delete/:id", handlers.DeletePost)

	e.GET("/", handlers.HomeHandler)
	e.GET("/detail/:id", handlers.DetailHandler)
	e.GET("/create", handlers.RenderPostPage, middleware.AuthMiddleware)
	e.POST("/login", handlers.HandleLogin)
	e.GET("/login", handlers.LoginPage)
	e.GET("/dashboard", handlers.DashboardViewHandler, middleware.AuthMiddleware)
	e.GET("/register", handlers.RegisterViewHandler)
	e.POST("/register", handlers.RegisterHandler)
}
