package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, echo.Map{
		"success" : true,
		"message" : message,
		"data" : data,
	})
}

func Created(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusCreated, echo.Map{
		"success" : true,
		"message" : message,
		"data" : data,
	})
}

func Error(c echo.Context, status int, message string) error {
	return c.JSON(status, echo.Map{
		"success" : false,
		"message" : message,
	})
}