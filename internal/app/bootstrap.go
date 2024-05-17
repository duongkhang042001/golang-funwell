package app

import (
	"core/pkg/csrf"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apiMiddlewares "core/internal/delivery/http/middleware"
	"core/internal/delivery/http/route"
)

func (app *Application) Bootstrap(e *echo.Echo) error {

	mw := apiMiddlewares.NewMiddlewareManager(app.config, []string{"*"}, app.logger)

	e.Use(mw.RequestLoggerMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RequestID())

	if app.config.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	userGroup := v1.Group("/users")
	authGroup := v1.Group("/auth")

	route.MapAuthRoutes(authGroup)
	route.MapUserRoutes(userGroup)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
