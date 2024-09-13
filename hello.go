package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Address string `json:"address" db:"address"`
}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

var (
	db              *sqlx.DB
	successResponse = Response{Message: "success executing the query", Status: true}
	errorResponse   = Response{Message: "failed executing the query", Status: false}
)

func initDB() error {
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres dbname=golang-test sslmode=disable host=localhost port=5432")
	if err != nil {
		return err
	}
	return db.Ping()
}

func getAllUsers(c echo.Context) error {
	var users []Employee
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := Employee{}
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	user := Employee{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	_, err := db.NamedExec("INSERT INTO users(name, phone, address) VALUES (:name, :phone, :address)", user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}
	return c.JSON(http.StatusOK, successResponse)
}

func updateUser(c echo.Context) error {
	user := Employee{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	user.Id = id
	_, err := db.NamedExec("UPDATE users SET name=:name, phone=:phone, address=:address WHERE id=:id", user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}
	return c.JSON(http.StatusOK, successResponse)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}
	return c.JSON(http.StatusOK, successResponse)
}

func main() {
	if err := initDB(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully Connected")

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/users", getAllUsers)
	e.GET("/user/:id", getUser)
	e.POST("/users", createUser)
	e.PUT("/users/update/:id", updateUser)
	e.DELETE("/users/delete/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8000"))
}
