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
	// Associate the /albums/:id path with the getAlbumByID function
	// : symbol in path signifies item is a path parameter
	router.GET("/albums/:id", getAlbumByID)
	// Use POST function to associate the POST HTTP method and /albums path with handler function
	router.POST("/albums", postAlbums)
	// Run function to attach the router to an http.Server and start the server
	router.Run("localhost:8080")
}

// getAlbums responds with a list of all albums as JSON
func getAlbums(c *gin.Context) {
	// gin.Context carries request details, validates/serializes JSON, and more

	// Serialize the struct into JSON and add to response
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJson to bind received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add new album struct from JSON to albums slice
	albums = append(albums, newAlbum)
	// Add 201 status code to response and JSON representing album created
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates album whose ID value matches the id parameter sent
// by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// Use Context.param to retrieve the id path parameter from the URL
	// When you map this handler to a path, you'll use a placeholder param in the path
	id := c.Param("id")

	// Loop over the list of albums looking for an album whose
	// ID value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			// 200 status code and serialize album struct and return in response as JSON
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// 404 not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
