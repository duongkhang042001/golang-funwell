package routes

import (
	"context"
	"core/pkg/validator"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=40"`
}

func InitUserRoutes(e *echo.Group) {

	users := e.Group("/users")

	users.GET(``, func(c echo.Context) error {
		user := &User{
			Username: "testuser",
			Email:    "testexample.com",
			Password: "pas",
		}

		err := validator.ValidateStruct(context.Background(), user)
		if err != nil {
			return c.JSON(400, map[string]string{"error": err.Error()})
		}

		return c.JSON(200, map[string]string{"status": "OK"})
	})
}
