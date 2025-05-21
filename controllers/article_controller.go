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
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	q := c.QueryParam("q")
	status := c.QueryParam("status")
	categoryID := c.QueryParam("category_id")
	tagIDs := c.QueryParam("tags") // format: uuid1,uuid2,...

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
	query := database.DB.Preload("Author").Preload("Category").Preload("Tags")

	// Filter status
	if status != "" {
		query = query.Where("status = ?", status)
	} else {
		query = query.Where("status = ?", "published")
	}

	// Search by title
	if q != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(q)+"%")
	}

	// Filter by category
	if categoryID != "" {
		if catUUID, err := uuid.Parse(categoryID); err == nil {
			query = query.Where("category_id = ?", catUUID)
		}
	}

	// Filter by tags
	if tagIDs != "" {
		tagIDList := strings.Split(tagIDs, ",")
		var tagUUIDs []uuid.UUID
		for _, tagID := range tagIDList {
			if tagUUID, err := uuid.Parse(strings.TrimSpace(tagID)); err == nil {
				tagUUIDs = append(tagUUIDs, tagUUID)
			}
		}
		if len(tagUUIDs) > 0 {
			query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
				Where("article_tags.tag_id IN ?", tagUUIDs).
				Group("articles.id")
		}
	}

	// Hitung total
	var total int64
	query.Model(&models.Article{}).Count(&total)

	// Ambil data dengan pagination
	if err := query.Offset(offset).Limit(limitInt).Find(&articles).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal mengambil data artikel")
	}

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
	if err := database.DB.Preload("Author").Preload("Category").Preload("Tags").
		Where("id = ? AND status = ?", id, "published").First(&article).Error; err != nil {
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
	categoryID := c.FormValue("category_id")
	tagIDs := c.FormValue("tags") // format: id1,id2,id3

	// Parse category ID
	catUUID, err := uuid.Parse(categoryID)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Kategori tidak valid")
	}

	// Parse tags
	var tagList []models.Tag
	if tagIDs != "" {
		tagIDStrs := strings.Split(tagIDs, ",")
		for _, idStr := range tagIDStrs {
			id, err := uuid.Parse(strings.TrimSpace(idStr))
			if err == nil {
				tagList = append(tagList, models.Tag{ID: id})
			}
		}
	}

	// Upload image
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
		Title:      title,
		Content:    content,
		Status:     status,
		ImageURL:   imageURL,
		AuthorID:   uuid.MustParse(userID),
		CategoryID: catUUID,
		Tags:       tagList,
	}

	if err := database.DB.Create(&article).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menyimpan artikel")
	}

	return response.Created(c, article, "Artikel berhasil dibuat")
}

// PUT /articles/:id
func UpdateArticle(c echo.Context) error {
	articleID := c.Param("id")
	id, err := uuid.Parse(articleID)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "ID artikel tidak valid")
	}

	var article models.Article
	if err := database.DB.Preload("Tags").Where("id = ?", id).First(&article).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}

	if !middleware.IsSelfOrAdmin(article.AuthorID, c) {
		return response.Error(c, http.StatusForbidden, "Anda tidak berhak mengakses artikel ini")
	}

	title := c.FormValue("title")
	content := c.FormValue("content")
	status := c.FormValue("status")
	categoryID := c.FormValue("category_id")
	tagIDs := c.FormValue("tags")

	if title != "" {
		article.Title = title
	}
	if content != "" {
		article.Content = content
	}
	if status != "" {
		article.Status = status
	}
	if categoryID != "" {
		if catUUID, err := uuid.Parse(categoryID); err == nil {
			article.CategoryID = catUUID
		}
	}
	if tagIDs != "" {
		var tagList []models.Tag
		for _, idStr := range strings.Split(tagIDs, ",") {
			id, err := uuid.Parse(strings.TrimSpace(idStr))
			if err == nil {
				tagList = append(tagList, models.Tag{ID: id})
			}
		}
		database.DB.Model(&article).Association("Tags").Replace(tagList)
	}

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
	if err := database.DB.Preload("Author").Preload("Category").Preload("Tags").
		Where("slug = ? AND status = ?", slug, "published").
		First(&article).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}

	return response.Success(c, article, "Detail artikel berhasil dimuat")
}
