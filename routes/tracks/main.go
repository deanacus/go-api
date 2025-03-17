package tracks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Track struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`

	Title    string `json:"title"`
	ArtistID int    `json:"artist_id"`
	AlbumID  int    `json:"album_id"`
	Length   int    `json:"length"`
}

var tracks = []Track{}

func InitRoutes(rg *gin.Engine) {
	tracks := rg.Group("/tracks")

	// create track list endpoint
	tracks.GET("/", listTracks)

	// create track create endpoint
	tracks.POST("/", createTrack)

	// create track search endpoint
	tracks.GET("/search", searchTracks)

	// create track details endpoint
	tracks.GET("/:id", getTrack)

	// create track update endpoint
	tracks.PATCH("/:id", updateTrack)
}

func listTracks(c *gin.Context) {
	c.JSON(http.StatusOK, tracks)
}

func createTrack(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func searchTracks(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func getTrack(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func updateTrack(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}
