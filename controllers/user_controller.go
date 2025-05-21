package controllers

import (
	"cms-go-2/config"
	"cms-go-2/database"
	"cms-go-2/models"
	"cms-go-2/response"
	"cms-go-2/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RegisterInput struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// POST /auth/register
func Register(c echo.Context) error {
	var input RegisterInput
	if err := c.Bind(&input); err != nil {
		return response.Error(c, http.StatusBadRequest, "Input tidak valid")
	}

	// Cari role default "operator"
	var role models.Role
	if err := database.DB.Where("name = ?", "operator").First(&role).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Role default tidak ditemukan")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal hash password")
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		RoleID:   role.ID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal registrasi user")
	}

	return response.Created(c, nil, "Registrasi berhasil")
}

// POST /auth/login
func Login(c echo.Context) error {
	var input LoginInput
	if err := c.Bind(&input); err != nil {
		return response.Error(c, http.StatusBadRequest, "Input tidak valid")
	}

	var user models.User
	if err := database.DB.Preload("Role").Where("email = ?", input.Email).First(&user).Error; err != nil {
		return response.Error(c, http.StatusUnauthorized, "Email tidak ditemukan")
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return response.Error(c, http.StatusUnauthorized, "Password salah")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"name":    user.Name,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := config.Get("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal generate token")
	}

	return response.Success(c, echo.Map{"token": tokenString}, "Login berhasil")
}

// POST /auth/refresh
func RefreshToken(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(jwt.MapClaims)

	// Buat token baru dengan data yang sama
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims["user_id"],
		"name":    claims["name"],
		"role":    claims["role"],
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := config.Get("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}

	tokenString, err := newToken.SignedString([]byte(secret))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal refresh token")
	}

	return response.Success(c, echo.Map{"token": tokenString}, "Token berhasil diperbarui")
}

// GET /user/me
func GetProfile(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(jwt.MapClaims)

	return response.Success(c, echo.Map{
		"user_id": claims["user_id"],
		"name":    claims["name"],
		"role":    claims["role"],
	}, "Profil pengguna berhasil dimuat")
}

// GET /admin/users
func GetAllUsers(c echo.Context) error {
	var users []models.User
	if err := database.DB.Preload("Role").Find(&users).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal mengambil data user")
	}
	return response.Success(c, users, "Data user berhasil dimuat")
}

// PUT /admin/users/:id/role
func UpdateUserRole(c echo.Context) error {
	userID := c.Param("id")
	roleID := c.FormValue("role_id")

	uid, err := uuid.Parse(userID)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "ID user tidak valid")
	}
	rid, err := uuid.Parse(roleID)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "ID role tidak valid")
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", uid).Error; err != nil {
		return response.Error(c, http.StatusNotFound, "User tidak ditemukan")
	}

	user.RoleID = rid
	if err := database.DB.Save(&user).Error; err != nil {
		return response.Error(c, http.StatusInternalServerError, "Gagal memperbarui role user")
	}

	return response.Success(c, user, "Role user berhasil diperbarui")
}
