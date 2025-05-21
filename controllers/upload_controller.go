package controllers

import (
	"cms-go-2/response"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "Gambar tidak ditemukan")
	}

	src, err := file.Open()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal membuka file")
	}
	defer src.Close()

	// Buat nama file unik
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	path := filepath.Join("uploads", filename)

	// Simpan file
	dst, err := os.Create(path)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menyimpan file")
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menyalin file")
	}

	// Buat URL (contoh lokal)
	imageURL := fmt.Sprintf("/uploads/%s", filename)

	return response.Success(c, echo.Map{
		"image_url": imageURL,
	}, "Upload berhasil")
}
