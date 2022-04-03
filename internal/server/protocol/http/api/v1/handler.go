package v1

import (
	"github.com/rs/zerolog"
)

type Handler struct {
	log zerolog.Logger
}

type option func(*Handler)

func NewHandler(opts ...option) *Handler {
	h := &Handler{}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func WithLogger(log zerolog.Logger) option {
	return func(h *Handler) {
		h.log = log
	}
}
