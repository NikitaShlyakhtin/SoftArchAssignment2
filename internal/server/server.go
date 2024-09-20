package server

import (
	"Assignment2/internal/app"
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.SugaredLogger
	echo   *echo.Echo
	app    *app.Application
}

func NewServer(app *app.Application) *Server {
	s := &Server{
		logger: app.Logger,
		app:    app,
		echo:   echo.New(),
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	s.echo.GET("/messages/count", s.app.GetMessageCount())
	s.echo.GET("/messages/ws", s.app.HandleWebSocket())
}

func (s *Server) Start(address string) error {
	s.logger.Infof("Starting server on address: %v", address)

	return s.echo.Start(address)
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server")

	return s.echo.Shutdown(ctx)
}
