package routes

import (
	"cms-go-2/controllers"
	"cms-go-2/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	auth.POST("/refresh", controllers.RefreshToken, middleware.JWTMiddleware) // ✅ DITAMBAHKAN

	user := e.Group("/user")
	user.Use(middleware.JWTMiddleware)
	user.GET("/me", controllers.GetProfile)
}

func UserManagerRoutes(e *echo.Echo) {
	admin := e.Group("/admin/users", middleware.JWTMiddleware, middleware.IsAdmin)
	admin.GET("", controllers.GetAllUsers)
	admin.PUT("/:id/role", controllers.UpdateUserRole)
}
