package middleware

import (
	"cms-go-2/response"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		if claims["role"] != "admin" {
			return response.Error(c, http.StatusForbidden, "Akses hanya untuk admin")
		}
		return next(c)
	}
}

func IsEditor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		role := claims["role"]
		if role != "editor" && role != "admin" {
			return response.Error(c, http.StatusForbidden, "Akses hanya untuk editor atau admin")
		}
		return next(c)
	}
}

func IsOperator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		if claims["role"] != "operator" {
			return response.Error(c, http.StatusForbidden, "Akses hanya untuk operator")
		}
		return next(c)
	}
}

// Mengecek apakah user adalah admin atau pemilik data (artikel, dsb)
func IsSelfOrAdmin(resourceAuthorID uuid.UUID, c echo.Context) bool {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	return role == "admin" || resourceAuthorID.String() == userID
}
