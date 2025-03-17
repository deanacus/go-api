package main

import (
	"github.com/deanacus/music-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.Init(server)
	server.Run(":8080")
}
