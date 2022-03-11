package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Initialize a Gin router using Default
	router := gin.Default()
	// Use the GET function to associate the GET HTTP method and /albums path with handler function
	router.GET("/albums", getAlbums)
	// Run function to attach the router to an http.Server and start the server
	router.Run("localhost:8080")
}

// getAlbums responds with a list of all albums as JSON
func getAlbums(c *gin.Context) {
	// gin.Context carries request details, validates/serializes JSON, and more

	// Serialize the struct into JSON and add to response
	c.IndentedJSON(http.StatusOK, albums)
}
