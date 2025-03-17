package routes

import (
	"github.com/deanacus/music-api/routes/albums"
	"github.com/deanacus/music-api/routes/artists"
	"github.com/deanacus/music-api/routes/scrobbles"
	"github.com/deanacus/music-api/routes/tracks"
	"github.com/deanacus/music-api/routes/users"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	albums.InitRoutes(router)
	artists.InitRoutes(router)
	tracks.InitRoutes(router)
	scrobbles.InitRoutes(router)
	users.InitRoutes(router)
}
