package main

import (
	"fmt"

	"github.com/2golang/vnexpress-rss/handlers/articles"
	"github.com/2golang/vnexpress-rss/middleware"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	api := r.Group("/v1")
	// public routes
	publicRoutes := api.Group("/")
	{
		publicRoutes.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	// protected routes
	protectedRoutes := api.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware())
	articleRoutes := protectedRoutes.Group("/articles")
	{
		articleRoutes.GET("", articles.ListArticles)
		articleRoutes.GET("/:id", articles.GetArticleById)
	}
}

func main() {
	// Creates a gin router with default middleware:
	r := gin.Default()
	// Setup routes
	setupRoutes(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"
	if err := r.Run(fmt.Sprintf(":%v", 8080)); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
