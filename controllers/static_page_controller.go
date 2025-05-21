package controllers

import (
	"cms-go-2/database"
	"cms-go-2/models"
	"cms-go-2/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStaticPage(c echo.Context) error {
	slug := c.Param("slug")

	var page models.StaticPage
	if err := database.DB.
		Where("slug = ? AND status = ?", slug, "published").
		First(&page).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Halaman tidak ditemukan atau belum dipublikasikan")
	}

	return response.Success(c, page, "Halaman berhasil dimuat")
}

// GET /admin/pages
func GetAllStaticPages(c echo.Context) error {
	var pages []models.StaticPage
	if err := database.DB.Find(&pages).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal memuat daftar halaman")
	}
	return response.Success(c, pages, "Daftar halaman berhasil dimuat")
}

// PUT /pages/:slug
func UpdateStaticPage(c echo.Context) error {
	slug := c.Param("slug")

	var page models.StaticPage
	if err := database.DB.Where("slug = ?", slug).First(&page).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Halaman tidak ditemukan")
	}

	// Update konten
	title := c.FormValue("title")
	content := c.FormValue("content")
	status := c.FormValue("status")

	if title != "" {
		page.Title = title
	}
	if content != "" {
		page.Content = content
	}
	if status != "" {
		page.Status = status
	}

	if err := database.DB.Save(&page).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal memperbarui halaman")
	}

	return response.Success(c, page, "Halaman berhasil diperbarui")
}
