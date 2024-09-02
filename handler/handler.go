package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanzohasashi17/restapi-guide/model"
)

// Get all albums
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Albums)
}

// Get by ID
func GetAlbumById (c *gin.Context) {
	id := c.Param("id")

	for _, a := range model.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, model.Error { Error: "Album by id: " + id + " is not found" })
}

// Create album
func PostAlbum(c *gin.Context) {
	var newAlbum model.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error { Error: "Invalid body"})
		return
	}
	model.Albums = append(model.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Delete album
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range model.Albums {
		if a.ID == id {
			model.Albums = append(model.Albums[:i], model.Albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H {"message": "Album by id: " + id + " is deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, nil)
}

// Update album
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	
	for i, a := range model.Albums {
		if a.ID == id {
			c.BindJSON(&a)
			model.Albums[i] = a
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, model.Error {"Not found"})
}
