package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`

	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{}

func InitRoutes(rg *gin.Engine) {
	users := rg.Group("/users")

	users.GET("/:id", getUser)

	// create user scrobbles endpoint
	users.GET("/:id/scrobbles", getScrobbles)

	// // create user artists endpoint
	// users.GET("/:id/artists", getArtists)

	// // create user albums endpoint
	// users.GET("/:id/albums", getAlbums)

	// // create user tracks endpoint
	// users.GET("/:id/tracks", getTracks)

	// // create user now-playing endpoint
	// users.GET("/:id/now-playing", getNowPlaying)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func getScrobbles(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
}

// func getArtists(c *gin.Context) {
// 	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
// }

// func getAlbums(c *gin.Context) {
// 	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
// }

// func getTracks(c *gin.Context) {
// 	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
// }

// func getNowPlaying(c *gin.Context) {
// 	c.JSON(http.StatusNotImplemented, gin.H{"message": "This endpoint has not yet been implemented"})
// }
