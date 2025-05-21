package controllers

import (
	"cms-go-2/database"
	"cms-go-2/models"
	"cms-go-2/response"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GET /articles
func GetArticles(c echo.Context) error {
	var articles []models.Article
	if err := database.DB.Preload("Author").Where("status = ?", "published").Find(&articles).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal mengambil data artikel")
	}
	return response.Success(c, articles, "Daftar artikel berhasil dimuat")
}

// GET /articles/:id
func GetArticleByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "ID tidak valid")
	}

	var article models.Article
	if err := database.DB.Preload("Author").Where("id = ? AND status = ?", id, "published").First(&article).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}
	return response.Success(c, article, "Detail artikel berhasil dimuat")
}

// POST /articles
func CreateArticle(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	var input models.Article
	if err := c.Bind(&input); err != nil {
		return response.Error(c, http.StatusBadRequest, "Input tidak valid")
	}

	input.AuthorID = uuid.MustParse(userID)

	if err := database.DB.Create(&input).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menyimpan artikel")
	}

	return response.Created(c, input, "Artikel berhasil dibuat")
}

// PUT /articles/:id
func UpdateArticle(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	articleID := c.Param("id")
	var article models.Article

	if err := database.DB.Where("id = ?", articleID).First(&article).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}

	// Hanya author yang boleh update
	if article.AuthorID.String() != userID {
		return response.Error(c, http.StatusForbidden, "Anda tidak berhak mengubah artikel ini")
	}

	var input models.Article
	if err := c.Bind(&input); err != nil {
		return response.Error(c, http.StatusBadRequest, "Input tidak valid")
	}

	article.Title = input.Title
	article.Content = input.Content
	article.Status = input.Status
	article.ImageURL = input.ImageURL

	if err := database.DB.Save(&article).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal mengupdate artikel")
	}

	return response.Success(c, article, "Artikel berhasil diperbarui")
}

// DELETE /articles/:id
func DeleteArticle(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	userRole := claims["role"].(string)

	articleID := c.Param("id")
	var article models.Article

	if err := database.DB.Where("id = ?", articleID).First(&article).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}

	// Admin boleh hapus semua, editor hanya milik sendiri
	if userRole != "admin" && article.AuthorID.String() != userID {
		return response.Error(c, http.StatusForbidden, "Anda tidak berhak menghapus artikel ini")
	}

	if err := database.DB.Delete(&article).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menghapus artikel")
	}

	return response.Success(c, nil, "Artikel berhasil dihapus")
}