package route

import (
	"core/internal/interfaces/http/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MapUserRoutes(router *echo.Group, mw *middleware.MiddlewareManager) {
	router.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Get all user page!")
	})
	router.GET("/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Get a user page with ID: "+c.Param("id"))
	})
	router.DELETE("/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Delete a user page with ID: "+c.Param("id"))
	})
	router.PUT("/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Update a user page with ID: "+c.Param("id"))
	})
	router.PATCH("/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Modify a user page with ID: "+c.Param("id"))
	})
	router.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, Create new user page!")
	})
}
