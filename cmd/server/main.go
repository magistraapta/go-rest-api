package main

import (
	"log"

	"example/hello/internal/api"
	"example/hello/internal/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully Connected")

	e := echo.New()
	e.Use(middleware.CORS())

	api.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
