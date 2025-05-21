package routes

import (
	"cms-go-2/controllers"

	"github.com/labstack/echo/v4"
)

func CategoryTagRoutes(e *echo.Echo) {
	e.GET("/categories", controllers.GetCategories)
	e.GET("/tags", controllers.GetTags)
}
