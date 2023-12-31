package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/v1shn3vsk7/cicd-lab/internal/app/http/handlers"
	"github.com/v1shn3vsk7/cicd-lab/internal/config"
	"github.com/v1shn3vsk7/cicd-lab/internal/tech/closer"
)

// Server ...
type Server struct {
	server   *echo.Echo
	handlers *handlers.Handlers
	cfg      *config.Config
}

// New ...
func New(cfg *config.Config) *Server {
	server := echo.New()
	server.Use(middleware.Recover())

	server.HideBanner = true
	server.HidePort = true

	return &Server{
		server:   server,
		handlers: handlers.New(),
		cfg:      cfg,
	}
}

// Start ...
func (s *Server) Start() {
	s.setupRoutes()

	go func() {
		log.Info().Msgf("starting listening http server at %s", s.cfg.HTTPAddr)
		err := s.server.Start(s.cfg.HTTPAddr)
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err)
		}
	}()

	closer.Add(s.close)
}

func (s *Server) close() error {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		log.Error().Msgf("error stopping http server, err: %v", err)
		return err
	}

	log.Info().Msgf("http server shutdown is done")

	return nil
}
