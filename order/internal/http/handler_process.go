package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type processRequest struct{}

type processResponse struct {
	Error string `json:"error,omitempty"`
}

func newProcessHandler(orderRepository orderRepository) *processHandler {
	return &processHandler{
		orderRepository: orderRepository,
	}
}

type processHandler struct {
	orderRepository orderRepository
}

func (h *processHandler) Handle(ctx echo.Context) error {
	var request getRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind create request: %v", err)
	}

	err = h.orderRepository.ProcessOrders()
	if err != nil {
		log.Warn(err)
		return ctx.JSON(http.StatusBadRequest, processResponse{Error: "cannot process orders"})
	}

	return ctx.JSON(http.StatusOK, processResponse{})
}
