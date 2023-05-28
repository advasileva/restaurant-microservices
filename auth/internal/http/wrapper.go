package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type errorResponse struct {
	Error string `json:"error,omitempty"`
}

type handler interface {
	Handle(ctx echo.Context) error
}

func newWrapper(handler handler) *wrapper {
	return &wrapper{
		wrapped: handler,
	}
}

type wrapper struct {
	wrapped handler
}

func (w *wrapper) Handle(ctx echo.Context) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Error(fmt.Errorf("panic: %v", err))
			ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "something goes wrong: panic occurred"})
		}
	}()

	err := w.wrapped.Handle(ctx)
	if err != nil {
		log.Error(fmt.Errorf("error: %v", err))
		return ctx.JSON(http.StatusInternalServerError, errorResponse{Error: "something goes wrong: error occurred"})
	}

	return nil
}
