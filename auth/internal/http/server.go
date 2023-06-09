package http

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewServer(
	userRepository userRepository,
	authService authService,
) (*server, error) {
	port, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse port: %v", err)
	}

	instance := echo.New()
	instance.Server.Addr = fmt.Sprintf(":%d", port)

	instance.Add("POST", "user/register", newWrapper(newRegisterHandler(userRepository)).Handle)
	instance.Add("POST", "user/auth", newWrapper(newAuthHandler(userRepository, authService)).Handle)
	instance.Add("GET", "user/info", newWrapper(newInfoHandler(userRepository, authService)).Handle)

	return &server{
		echo: instance,
	}, nil
}

type server struct {
	echo *echo.Echo
}

func (s *server) Serve() error {
	return s.echo.StartServer(s.echo.Server)
}
