package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong pong",
	})
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
