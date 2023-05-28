package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userRegistrar interface {
	Register() error
}

func newRegisterHandler(userRegisterer userRegistrar) *handler {
	return &handler{
		userRegistrar: userRegisterer,
	}
}

type handler struct {
	userRegistrar userRegistrar
}

func (h *handler) Handle(ctx echo.Context) error {
	var request registerRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind register request: %v", err)
	}

	return ctx.JSON(http.StatusOK, registerResponse{
		Smth: request.Smth + 1,
	})
}
