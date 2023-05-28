package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userInfoRequest struct{}

type userInfoResponse struct {
	Username string `json:"username"`
	Error    string `json:"error,omitempty"`
}

func newUserInfoHandler(userInformer userRepository, authService authService) *userInfoHandler {
	return &userInfoHandler{
		userInformer: userInformer,
		authService:  authService,
	}
}

type userInfoHandler struct {
	userInformer userRepository
	authService  authService
}

func (h *userInfoHandler) Handle(ctx echo.Context) error {
	var request userInfoRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind user_info request: %v", err)
	}

	header := ctx.Request().Header.Get("Authorization")
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, userInfoResponse{Error: "empty Authorization header"})
	}

	parts := strings.Split(header, " ")
	if len(parts) < 3 {
		return ctx.JSON(http.StatusUnauthorized, userInfoResponse{Error: "incorrect Authorization header"})
	}

	email, err := h.authService.GetUserEmailByJWT(parts[3])
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, userInfoResponse{Error: "incorrect token"})
	}

	username, err := h.userInformer.GetUsernameByEmail(email)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, userInfoResponse{Error: "username not found"})
	}

	return ctx.JSON(http.StatusOK, userInfoResponse{
		Username: username,
	})
}
