package routes

import (
	"cms-go-2/controllers"
	"cms-go-2/middleware"

	"github.com/labstack/echo/v4"
)

func ArticleRoutes(e *echo.Echo) {
	article := e.Group("/articles")
	article.GET("", controllers.GetArticles)
	article.GET("/:id", controllers.GetArticleByID)
	article.GET("/slug/:slug", controllers.GetArticleBySlug)

	// proteksi dengan JWT
	article.Use(middleware.JWTMiddleware)
	article.POST("", controllers.CreateArticle)
	article.PUT("/:id", controllers.UpdateArticle)
	article.DELETE("/:id", controllers.DeleteArticle)
}
