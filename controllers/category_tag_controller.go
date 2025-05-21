package controllers

import (
	"cms-go-2/database"
	"cms-go-2/models"
	"cms-go-2/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := database.DB.Order("name").Find(&categories).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal memuat kategori")
	}
	return response.Success(c, categories, "Kategori berhasil dimuat")
}

func GetTags(c echo.Context) error {
	var tags []models.Tag
	if err := database.DB.Order("name").Find(&tags).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal memuat tag")
	}
	return response.Success(c, tags, "Tag berhasil dimuat")
}
