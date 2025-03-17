package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`

	Title    string `json:"title"`
	ArtistID string `json:"artist_id"`
	Type     string `json:"type"`
	Year     string `json:"year"`
}

var albums = []Album{}

func InitRoutes(router *gin.Engine) {
	albums := router.Group("/albums")

	albums.GET("/", listAlbums)
	albums.POST("/", createAlbum)

	albums.GET("/search", searchAlbums)

	albums.GET("/:id", getAlbum)
	albums.GET("/:id/tracks", getAlbumTracks)
	albums.PATCH("/:id", updateAlbum)
}

func listAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func searchAlbums(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

// Get a single album
func getAlbum(c *gin.Context) {
	id := c.Param("id")

	// Iterate over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			// We found a match, so return it
			c.JSON(http.StatusOK, album)
			return
		}
	}

	// We didn't find a match, so return a 404 HTTP status
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbumTracks(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

// Create a new album
func createAlbum(c *gin.Context) {
	// Create a new album to populate with data from the request body
	var newAlbum Album

	// Bind the JSON to the new album.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
}

// Update an album TODO: refactor to do a PATCH instead of PUT (updateAlbum instead of replace)
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	// Create a new album to hold the updated values.
	var updatedAlbum Album

	// Bind the JSON to the new album.
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	// Iterate over the existing albums, and if we find
	// an album with a matching ID, update the album
	// with the new values.
	for i, album := range albums {
		if album.ID == id {
			albums[i] = updatedAlbum
			c.JSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	// If we didn't find an album with a matching ID,
	// return a 404 not found status.
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
