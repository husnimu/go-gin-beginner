package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Husni Mubarok",
		"bio":  "app dev",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "test 123",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param(("title"))
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	genre := c.Query("genre")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"genre": genre,
	})
}
