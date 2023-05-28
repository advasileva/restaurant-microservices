package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type infoRequest struct{}

type infoResponse struct {
	Username string `json:"username"`
	Error    string `json:"error,omitempty"`
}

func newInfoHandler(userRepository userRepository, authService authService) *infoHandler {
	return &infoHandler{
		userRepository: userRepository,
		authService:    authService,
	}
}

type infoHandler struct {
	userRepository userRepository
	authService    authService
}

func (h *infoHandler) Handle(ctx echo.Context) error {
	var request infoRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind info request: %v", err)
	}

	header := ctx.Request().Header.Get("Authorization")
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, infoResponse{Error: "empty Authorization header"})
	}

	parts := strings.Split(header, " ")
	if len(parts) < 3 {
		return ctx.JSON(http.StatusUnauthorized, infoResponse{Error: "incorrect Authorization header"})
	}

	email, err := h.authService.GetUserEmailByJWT(parts[2])
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, infoResponse{Error: "incorrect token"})
	}

	username, err := h.userRepository.GetUsernameByEmail(email)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, infoResponse{Error: "username not found"})
	}

	return ctx.JSON(http.StatusOK, infoResponse{
		Username: username,
	})
}
