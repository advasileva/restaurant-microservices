package http

import (
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type createRequest struct {
	UserId          int64         `json:"user_id"`
	Dishes          []models.Dish `json:"dishes"`
	Status          string        `json:"status"`
	SpecialRequests string        `json:"special_requests"`
}

type createResponse struct {
	Error string `json:"error,omitempty"`
}

func newCreateHandler(orderRepository orderRepository) *createHandler {
	return &createHandler{
		orderRepository: orderRepository,
	}
}

type createHandler struct {
	orderRepository orderRepository
}

func (h *createHandler) Handle(ctx echo.Context) error {
	var request createRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind create request: %v", err)
	}

	err = h.orderRepository.AddOrder(
		request.UserId,
		request.Dishes,
		request.Status,
		request.SpecialRequests,
	)
	if err != nil {
		log.Warn(err)
		return ctx.JSON(http.StatusBadRequest, createResponse{Error: "cannot create order"})
	}

	return ctx.JSON(http.StatusOK, createResponse{})
}
