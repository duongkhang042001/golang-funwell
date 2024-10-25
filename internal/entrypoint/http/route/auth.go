package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(router *echo.Group) {
	router.POST("/sign-in", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Sign in page!")
	})

	router.GET("/sign-in", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Sign in page!")
	})
	router.POST("/sign-up", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Sign up page!")
	})
}
