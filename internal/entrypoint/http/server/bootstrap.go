package server

import (
	"core/pkg/csrf"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apiMiddlewares "core/internal/entrypoint/http/middleware"
	"core/internal/entrypoint/http/route"
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
	//TODO: Init repositories
	//TODO: Init useCases
	v1 := e.Group("/api/v1")

	route.MapAuthRoutes(v1.Group("/auth"))

	health := v1.Group("/health")

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}

// DOCS: https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API/blob/master/internal/server/handlers.go
