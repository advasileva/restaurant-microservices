package http

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewServer(
	orderRepository orderRepository,
) (*server, error) {
	port, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse port: %v", err)
	}

	instance := echo.New()
	instance.Server.Addr = fmt.Sprintf(":%d", port)

	instance.Add("POST", "order/create", newWrapper(newCreateHandler(orderRepository)).Handle)
	instance.Add("POST", "order/process", newWrapper(newProcessHandler(orderRepository)).Handle)
	instance.Add("GET", "order/:orderId", newWrapper(newGetHandler(orderRepository)).Handle)

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
