package server

import (
	"context"
	"core/config"
	"core/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	certFile       = "ssl/Server.crt"
	keyFile        = "ssl/Server.pem"
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Application struct {
	echo   *echo.Echo
	config *config.Config
	logger logger.Logger
}

func NewApplication(config *config.Config, logger logger.Logger) *Application {
	return &Application{
		echo:   echo.New(),
		config: config,
		logger: logger,
	}
}

func (app *Application) Start() error {

	if app.config.Server.SSL {
		if err := app.Bootstrap(app.echo); err != nil {
			return err
		}

		app.echo.Server.ReadTimeout = time.Second * app.config.Server.ReadTimeout
		app.echo.Server.WriteTimeout = time.Second * app.config.Server.WriteTimeout

		go func() {
			app.logger.Infof("Server is listening on PORT: %s", app.config.Server.Port)
			app.echo.Server.ReadTimeout = time.Second * app.config.Server.ReadTimeout
			app.echo.Server.WriteTimeout = time.Second * app.config.Server.WriteTimeout
			app.echo.Server.MaxHeaderBytes = maxHeaderBytes
			if err := app.echo.StartTLS(app.config.Server.Port, certFile, keyFile); err != nil {
				app.logger.Fatalf("Error starting TLS Server: ", err)
			}
		}()

		go func() {
			app.logger.Infof("Starting Debug Server on PORT: %s", app.config.Server.PprofPort)
			if err := http.ListenAndServe(app.config.Server.PprofPort, http.DefaultServeMux); err != nil {
				app.logger.Errorf("Error PPROF ListenAndServe: %s", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		<-quit

		ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
		defer shutdown()

		app.logger.Info("Server Exited Properly")
		return app.echo.Server.Shutdown(ctx)
	}

	server := &http.Server{
		Addr:           app.config.Server.Port,
		ReadTimeout:    time.Second * app.config.Server.ReadTimeout,
		WriteTimeout:   time.Second * app.config.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		app.logger.Infof("Server is listening on PORT: %s", app.config.Server.Port)
		if err := app.echo.StartServer(server); err != nil {
			app.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	go func() {
		app.logger.Infof("Starting Debug Server on PORT: %s", app.config.Server.PprofPort)
		if err := http.ListenAndServe(app.config.Server.PprofPort, http.DefaultServeMux); err != nil {
			app.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	if err := app.Bootstrap(app.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	app.logger.Info("Server Exited Properly")

	return app.echo.Server.Shutdown(ctx)
}
