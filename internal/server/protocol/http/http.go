package http

import (
	v1 "github.com/VyalovAlexander/Mytheresa/internal/server/protocol/http/api/v1"
	"net/http"

	"github.com/labstack/echo/v4"

)

type Config struct {
	Addr  string
}

type Server struct {
	*echo.Echo
	*v1.Handler
	swaggerAssets     http.Handler
}

type option func(*Server)

func WithAddr(addr string) option {
	return func(s *Server) {
		s.Server.Addr = addr
	}
}

func WithLogger(log echo.Logger) option {
	return func(s *Server) {
		s.Logger = log
	}
}

func WithHandler(h *v1.Handler) option {
	return func(s *Server) {
		s.Handler = h
	}
}

// nolint funlen
func NewServer(opts ...option) *Server {
	srv := &Server{}
	srv.Echo = echo.New()

	srv.HideBanner = true
	srv.HidePort = true

	for _, opt := range opts {
		opt(srv)
	}
	api := srv.Group("/api")
	v1Group := api.Group("/v1")
	v1Group.GET("/products", srv.Handler.Products)

	return srv
}

func (s *Server) Serve() error {
	s.Logger.Infof("http server started on %s", s.Server.Addr)
	return s.Start(s.Server.Addr)
}