package main

import (
	"cms-go-2/config"
	"cms-go-2/database"
	"cms-go-2/response"
	"cms-go-2/routes"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load konfigurasi dari .env
	config.LoadEnv()

	// Ambil port dari .env
	port := config.Get("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Inisialisasi koneksi database
	err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// Inisialisasi Echo
	e := echo.New()
	e.Static("/uploads", "uploads")


	// Tes route awal
	e.GET("/", func(c echo.Context) error {
		return response.Success(c, "CMS-Go API is running properly.", "Server siap")
	})

	// Inisialisasi routes
	routes.UserRoutes(e)
	routes.ArticleRoutes(e)
	routes.UploadRoutes(e)
	routes.StaticPageRoutes(e)
	routes.CategoryTagRoutes(e)

	// Jalankan server
	fmt.Println("Server berjalan di port:", port)
	e.Logger.Fatal(e.Start(":" + port))
}
