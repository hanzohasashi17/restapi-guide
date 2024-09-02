package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hanzohasashi17/restapi-guide/handler"
)

func main() {
	router := gin.Default()

	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumById)
	router.POST("/albums", handler.PostAlbum)
	router.DELETE("/albums/:id", handler.DeleteAlbum)
	router.PUT("/albums/:id", handler.UpdateAlbum)

	router.Run("localhost:8080")
}

