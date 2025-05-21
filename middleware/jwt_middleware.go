package middleware

import (
	"cms-go-2/config"
	"cms-go-2/response"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, http.StatusUnauthorized, "Token tidak ditemukan")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		secret := config.Get("JWT_SECRET")
		if secret == "" {
			secret = "secret"
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Metode signing tidak valid")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return response.Error(c, http.StatusUnauthorized, "Token tidak valid")
		}

		// Set token di context agar bisa diakses di controller
		c.Set("user", token)

		return next(c)
	}
}
