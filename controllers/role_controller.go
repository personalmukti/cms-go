package controllers

import (
	"cms-go-2/database"
	"cms-go-2/models"
	"cms-go-2/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GET /admin/roles
func GetAllRoles(c echo.Context) error {
	var roles []models.Role
	if err := database.DB.Find(&roles).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal memuat data role")
	}
	return response.Success(c, roles, "Data role berhasil dimuat")
}

// POST /admin/roles
func CreateRole(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return response.Error(c, http.StatusBadRequest, "Nama role tidak boleh kosong")
	}

	role := models.Role{
		ID:   uuid.New(),
		Name: name,
	}
	if err := database.DB.Create(&role).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menambahkan role")
	}
	return response.Success(c, role, "Role berhasil ditambahkan")
}

// PUT /admin/roles/:id
func UpdateRole(c echo.Context) error {
	id := c.Param("id")
	newName := c.FormValue("name")

	rid, err := uuid.Parse(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "ID role tidak valid")
	}

	var role models.Role
	if err := database.DB.First(&role, "id = ?", rid).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "Role tidak ditemukan")
	}

	role.Name = newName
	if err := database.DB.Save(&role).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal memperbarui role")
	}
	return response.Success(c, role, "Role berhasil diperbarui")
}

// DELETE /admin/roles/:id
func DeleteRole(c echo.Context) error {
	id := c.Param("id")
	rid, err := uuid.Parse(id)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "ID role tidak valid")
	}

	if err := database.DB.Delete(&models.Role{}, "id = ?", rid).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal menghapus role")
	}
	return response.Success(c, nil, "Role berhasil dihapus")
}
