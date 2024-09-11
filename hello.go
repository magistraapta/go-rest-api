package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

type album struct {
	ID     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

var albums = []album{
	{ID: "1", Title: "whats up story morning glory", Artist: "Oasis", Price: 56.99},
	{ID: "2", Title: "for all the dogs", Artist: "Drake", Price: 100},
	{ID: "3", Title: "menari dalam bayangan", Artist: "Hindia", Price: 60.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/albums", postAlbum)
	router.Run("localhost:8000")
}
