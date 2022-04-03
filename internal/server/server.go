package server

import (
	"context"
	"errors"
)

const (
	ProtoHTTP = "http"
)

const (
	APIVerV3 = "v1"
)

var (
	ErrProtoUnknown  = errors.New("server: proto unknown")
	ErrAPIVerUnknown = errors.New("server: api version unknown")
)

type SMS struct {
	Bind struct{}
}

type Config struct {
	Bind     string           `mapstructure:"BIND" valid:"required"`
	APIVer   string           `mapstructure:"API_VERSION" valid:"required"`
}

type Server interface {
	Serve() error
	Shutdown(ctx context.Context) error
}
