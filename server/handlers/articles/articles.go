package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListArticles godoc
// @Summary List all articles
// @Description Get a list of all available articles
// @Tags articles
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{}
// @Router /articles [get]
func ListArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong pong",
	})
}

// GetArticleById godoc
// @Summary Get article by ID
// @Description Get article details by its ID
// @Tags articles
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Article ID"
// @Success 200 {object} map[string]interface{}
// @Router /articles/{id} [get]
func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
