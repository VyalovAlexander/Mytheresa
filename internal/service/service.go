package service

import (
	"context"
	"github.com/VyalovAlexander/Mytheresa/internal/server"
	"github.com/rs/zerolog"
)

type Service struct {
	log  zerolog.Logger
	srv  server.Server     `valid:"required"`
}

type option func(*Service)

func New(opts ...option) *Service {
	svc := &Service{}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}


func WithLogger(log zerolog.Logger) option {
	return func(s *Service) {
		s.log = log
	}
}

func WithServer(srv server.Server) option {
	return func(s *Service) {
		s.srv = srv
	}
}

func (s *Service) Start() error {
	s.log.Info().Msg("service started")
	return s.srv.Serve()
}

func (s *Service) Shutdown(ctx context.Context) error {
	s.log.Info().Msg("shutting down the service")
	return s.srv.Shutdown(ctx)
}
