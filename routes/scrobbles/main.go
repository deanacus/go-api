package scrobbles

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Scrobble struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`

	AlbumID    string `json:"album_id"`
	ArtistID   string `json:"artist_id"`
	ListenedAt string `json:"listened_at"`
	TrackID    string `json:"track_id"`
	UserID     string `json:"user_id"`
}

var scrobbles = []Scrobble{}

func InitRoutes(router *gin.Engine) {
	scrobbles := router.Group("/scrobbles")

	scrobbles.GET("/", listScrobbles)
	scrobbles.POST("/", createScrobble)
}

func listScrobbles(c *gin.Context) {
	userID, success := c.GetQuery("user")

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}

	fmt.Printf("Getting scrobbles for user %s", userID)
	foundedScrobbles := []Scrobble{}
	for _, scrobble := range scrobbles {
		if scrobble.UserID == userID {
			foundedScrobbles = append(foundedScrobbles, scrobble)
		}
	}
	c.JSON(http.StatusOK, foundedScrobbles)
}

func createScrobble(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}
