package routes

import "github.com/labstack/echo/v4"

func InitUserRoutes(e *echo.Group) {
	e.GET(`/users`, func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
}
