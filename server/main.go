package main

import (
	"fmt"

	docs "github.com/2golang/vnexpress-rss/docs"
	"github.com/2golang/vnexpress-rss/handlers/articles"
	"github.com/2golang/vnexpress-rss/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           VnExpress RSS API
// @version         1.0
// @description     This is api documentation for VnExpress RSS API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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
	// api docs
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
