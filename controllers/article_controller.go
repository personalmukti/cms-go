package controllers

import (
	"cms-go-2/database"
	"cms-go-2/middleware"
	"cms-go-2/models"
	"cms-go-2/response"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GET /articles
func GetArticles(c echo.Context) error {
	// Query parameter
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	q := c.QueryParam("q")
	status := c.QueryParam("status")

	// Default nilai
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	if pageInt <= 0 {
		pageInt = 1
	}
	if limitInt <= 0 {
		limitInt = 10
	}
	offset := (pageInt - 1) * limitInt

	var articles []models.Article
	query := database.DB.Preload("Author")

	// Filter status
	if status != "" {
		query = query.Where("status = ?", status)
	} else {
		query = query.Where("status = ?", "published")
	}

	// Search title
	if q != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(q)+"%")
	}

	// Hitung total
	var total int64
	query.Model(&models.Article{}).Count(&total)

	// Ambil data dengan pagination
	if err := query.Offset(offset).Limit(limitInt).Find(&articles).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal mengambil data artikel")
	}

	// Format respon paginasi
	result := echo.Map{
		"data":       articles,
		"page":       pageInt,
		"limit":      limitInt,
		"total":      total,
		"total_page": int((total + int64(limitInt) - 1) / int64(limitInt)),
	}

	return response.Success(c, result, "Daftar artikel berhasil dimuat")
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

	title := c.FormValue("title")
	content := c.FormValue("content")
	status := c.FormValue("status")

	// Upload file gambar jika ada
	imageURL := ""
	file, err := c.FormFile("image")
	if err == nil {
		src, _ := file.Open()
		defer src.Close()

		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		dstPath := filepath.Join("uploads", filename)
		dst, _ := os.Create(dstPath)
		defer dst.Close()

		io.Copy(dst, src)
		imageURL = "/uploads/" + filename
	}

	article := models.Article{
		Title:     title,
		Content:   content,
		Status:    status,
		ImageURL:  imageURL,
		AuthorID:  uuid.MustParse(userID),
	}

	if err := database.DB.Create(&article).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menyimpan artikel")
	}

	return response.Created(c, article, "Artikel berhasil dibuat")
}


// PUT /articles/:id
func UpdateArticle(c echo.Context) error {
	articleID := c.Param("id")

	// Validasi UUID
	id, err := uuid.Parse(articleID)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "ID artikel tidak valid")
	}

	var article models.Article
	if err := database.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}

	if !middleware.IsSelfOrAdmin(article.AuthorID, c) {
		return response.Error(c, http.StatusForbidden, "Anda tidak berhak mengakses artikel ini")
	}

	// Ambil data baru dari form
	title := c.FormValue("title")
	content := c.FormValue("content")
	status := c.FormValue("status")

	if title != "" {
		article.Title = title
	}
	if content != "" {
		article.Content = content
	}
	if status != "" {
		article.Status = status
	}

	// Jika ada file gambar baru, ganti
	file, err := c.FormFile("image")
	if err == nil {
		src, _ := file.Open()
		defer src.Close()

		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		dstPath := filepath.Join("uploads", filename)
		dst, _ := os.Create(dstPath)
		defer dst.Close()

		io.Copy(dst, src)
		article.ImageURL = "/uploads/" + filename
	}

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
// GET /articles/slug/:slug
func GetArticleBySlug(c echo.Context) error {
	slug := c.Param("slug")

	var article models.Article
	if err := database.DB.Preload("Author").
		Where("slug = ? AND status = ?", slug, "published").
		First(&article).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}

	return response.Success(c, article, "Detail artikel berhasil dimuat")
}
