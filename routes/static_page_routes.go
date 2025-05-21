package routes

import (
	"cms-go-2/controllers"
	"cms-go-2/middleware"

	"github.com/labstack/echo/v4"
)

func StaticPageRoutes(e *echo.Echo) {
	// Endpoint publik (tidak butuh token)
	e.GET("/pages/:slug", controllers.GetStaticPage)

	// Endpoint proteksi: hanya admin/operator (via IsContentManager)
	admin := e.Group("/admin/pages", middleware.JWTMiddleware, middleware.IsContentManager)
	admin.GET("", controllers.GetAllStaticPages)

	secured := e.Group("/pages", middleware.JWTMiddleware, middleware.IsContentManager)
	secured.PUT("/:slug", controllers.UpdateStaticPage)
}
