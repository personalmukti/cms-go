package routes

import (
	"cms-go-2/controllers"
	"cms-go-2/middleware"

	"github.com/labstack/echo/v4"
)

func RoleManagerRoutes(e *echo.Echo) {
	admin := e.Group("/admin/roles", middleware.JWTMiddleware, middleware.IsAdmin)
	admin.GET("", controllers.GetAllRoles)
	admin.POST("", controllers.CreateRole)
	admin.PUT("/:id", controllers.UpdateRole)
	admin.DELETE("/:id", controllers.DeleteRole)
}
