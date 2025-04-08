package artists

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model 	"github.com/deanacus/music-api/models/artists"

)

type Artist struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`

	Name string `json:"name"`
	Type string `json:"type"`
}

var artists = []Artist{}

func InitRoutes(router *gin.Engine) {
	artists := router.Group("/artists")

	// create artist list endpoint
	artists.GET("/", listArtists)

	// create artist create endpoint
	artists.POST("/", createArtist)

	// create artist search endpoint
	artists.GET("/search", searchArtists)

	// create artist details endpoint
	artists.GET("/:id", getArtist)
	// create artist update endpoint
	artists.PATCH("/:id", updateArtist)

	// create artist albums endpoint
	artists.GET("/:id/albums", getArtistAlbums)

	// create artist tracks endpoint
	artists.GET("/:id/tracks", getArtistTracks)

}

func listArtists(c *gin.Context) {
	model.getAlbum()
	c.JSON(http.StatusOK, artists)
}

func createArtist(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func searchArtists(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func getArtist(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func updateArtist(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func getArtistAlbums(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

func getArtistTracks(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}
