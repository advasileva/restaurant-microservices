package http

import (
	"fmt"
	"net/http"
	"server/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getRequest struct{}

type getResponse struct {
	models.Order
	Error string `json:"error,omitempty"`
}

func newGetHandler(orderRepository orderRepository) *getHandler {
	return &getHandler{
		orderRepository: orderRepository,
	}
}

type getHandler struct {
	orderRepository orderRepository
}

func (h *getHandler) Handle(ctx echo.Context) error {
	var request getRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind create request: %v", err)
	}

	id, err := strconv.ParseInt(ctx.Param("orderId"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, getResponse{Error: "cannot parse orderId"})
	}

	order, err := h.orderRepository.GetOrder(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, getResponse{Error: "cannot get order"})
	}

	return ctx.JSON(http.StatusOK, getResponse{Order: order})
}
