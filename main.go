package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Album struct {

	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{

	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbums)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func addAlbums(c *gin.Context) {
	var newAlbum Album
	err := c.BindJSON(&newAlbum); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid data"})
		return 
	}
	
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}