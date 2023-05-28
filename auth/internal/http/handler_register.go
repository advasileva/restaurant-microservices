package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var roles = map[string]bool{
	"customer": true,
	"chef":     true,
	"manager":  true,
}

type registerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	Error string `json:"error,omitempty"`
}

func newRegisterHandler(userRegisterer userRepository) *registerHandler {
	return &registerHandler{
		userRegistrar: userRegisterer,
	}
}

type registerHandler struct {
	userRegistrar userRepository
}

func (h *registerHandler) Handle(ctx echo.Context) error {
	var request registerRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind register request: %v", err)
	}

	if !h.isCorrectEmail(request.Email) {
		return ctx.JSON(http.StatusBadRequest, registerResponse{Error: "incorrect email"})
	}

	role := ctx.QueryParam("role")
	if !h.isCorrectRole(role) {
		return ctx.JSON(http.StatusBadRequest, registerResponse{Error: "incorrect role"})
	}

	err = h.userRegistrar.Add(
		request.Username,
		request.Email,
		request.Password,
		role,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, registerResponse{Error: "cannot register user"})
	}

	return ctx.JSON(http.StatusOK, registerResponse{})
}

func (h *registerHandler) isCorrectEmail(email string) bool {
	return strings.Contains(email, "@")
}

func (h *registerHandler) isCorrectRole(role string) bool {
	_, ok := roles[role]
	return ok
}
