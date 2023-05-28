package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

func newAuthHandler(userAuthentificator userRepository, authService authService) *authHandler {
	return &authHandler{
		userAuthentificator: userAuthentificator,
		authService:         authService,
	}
}

type authHandler struct {
	userAuthentificator userRepository
	authService         authService
}

func (h *authHandler) Handle(ctx echo.Context) error {
	var request authRequest
	err := ctx.Bind(&request)
	if err != nil {
		return fmt.Errorf("cannot bind auth request: %v", err)
	}

	ok, err := h.userAuthentificator.CheckPasswordByEmail(request.Email, request.Password)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, authResponse{Error: "incorrect password for email"})
	}

	token, err := h.authService.GetUserJWT(request.Email)
	if err != nil {
		return fmt.Errorf("cannot create jwt token: %v", err)
	}

	return ctx.JSON(http.StatusOK, authResponse{
		Token: token,
	})
}
