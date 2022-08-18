// Package utils providers utils
package utils

import (
	"github.com/apex/log"
	"lolesports/infrastructure/application"
)

// Handler implementation.
type Handler struct {
	Handlers []log.Handler
}

// New handler.
func New(h ...log.Handler) *Handler {
	return &Handler{
		Handlers: h,
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	for _, handler := range h.Handlers {
		if err := handler.HandleLog(e); err != nil {
			return err
		}
	}

	return nil
}

func BaseLogCtx() *log.Entry {
	ctx := &log.Entry{}
	ctx.Fields = log.Fields{
		"app":  "scrap",
		"env":  application.Context().Environment(),
		"role": application.Context().Environment(),
	}
	return ctx
}

func WithFields(ctx *log.Entry, fields log.Fields) {
	for key, msg := range fields {
		ctx.Fields[key] = msg
	}
}
