package routes

import (
	"cms-go-2/controllers"

	"github.com/labstack/echo/v4"
)

func UploadRoutes(e *echo.Echo) {
	e.POST("/upload", controllers.UploadImage)
}
