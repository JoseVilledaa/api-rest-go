package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Alone at prom", Artist: "Tory Lanez", Year: 2022},
	{ID: "2", Title: "Flower Boy", Artist: "Tyler, the creator", Year: 2017},
	{ID: "3", Title: "The Divine Feminine", Artist: "Mac Miller", Year: 2016},
}

// Get All
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Get By ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func createAlbum(c *gin.Context) {
	var newAlbum album

	c.BindJSON(&newAlbum)

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", createAlbum)
	router.Run("localhost:8080")
}
