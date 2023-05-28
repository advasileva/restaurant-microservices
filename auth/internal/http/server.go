package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

func NewServer(
	userRepository userRepository,
) (*server, error) {
	port, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse port: %v", err)
	}

	echo := echo.New()
	echo.Server.Addr = fmt.Sprintf(":%d", port)

	echo.Add("GET", "register", newRegisterHandler(userRepository).Handle)

	return &server{
		echo: echo,
	}, nil
}

type server struct {
	echo *echo.Echo
}

func (s *server) Serve() error {
	return s.echo.StartServer(s.echo.Server)
}
