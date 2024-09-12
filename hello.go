package main

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"

	sqlx "github.com/jmoiron/sqlx" //make alias name the package to sqlx
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq" //save it into underscore variable
)

type Employee struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Address string `json:"address" db:"address"`
}

type Response struct {
	Message string
	Status  bool
}

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=golang-test sslmode=disable host=localhost port=5432")

	if err != nil {
		log.Fatalln(err)
	}

	response := Response{
		Message: "success executing the query",
		Status:  true,
	}

	responseError := Response{
		Message: "failed executing the query",
		Status:  false,
	}

	e := echo.New()
	e.Use(middleware.CORS())

	// get all users
	e.GET("/users", func(c echo.Context) error {
		rows, _ := db.Queryx("select * from users")

		var users []Employee

		for rows.Next() {
			place := Employee{}
			rows.StructScan(&place)
			users = append(users, place)
		}

		return c.JSON(http.StatusOK, users)
	})

	// get spesific user
	e.GET("/user/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		user := Employee{}

		err = db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseError)
		}

		return c.JSON(http.StatusOK, user)
	})

	e.POST("/users", func(c echo.Context) error {
		reqBody := Employee{}
		c.Bind(&reqBody)

		_, err = db.NamedExec("INSERT INTO users(name, phone, address) VALUES (:name, :phone, :address)", reqBody)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseError)
		}

		return c.JSON(http.StatusOK, response)
	})

	e.PUT("/users/update/:id", func(c echo.Context) error {
		reqBody := Employee{}

		if err := c.Bind(&reqBody); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		id, _ := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		}

		reqBody.Id = id

		_, errQuery := db.NamedExec("update users SET name= :name, phone= :phone, address= :address WHERE id= :id", reqBody)

		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, responseError)
		}

		return c.JSON(http.StatusOK, response)
	})

	e.DELETE("/users/delete/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err = db.Exec("DELETE FROM users WHERE id = $1", id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseError)
		}
		return c.JSON(http.StatusOK, response)
	})

	// check database connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	//start server at port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
