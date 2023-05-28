package main

import (
	"fmt"
	"server/internal/database"
	"server/internal/http"
	"server/internal/repositories/order"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
)

func main() {
	db := database.Connect()
	defer func(db *pg.DB) {
		err := database.Disconnect(db)
		if err != nil {
			log.Error(fmt.Errorf("cannot disconnect with db: %v", err))
		}
	}(db)

	orderRepository := order.NewRepository(db)
	err := orderRepository.SetupTable()
	if err != nil {
		log.Error(fmt.Errorf("cannot setup table in order repository: %v", err))
		return
	}

	server, err := http.NewServer(orderRepository)
	if err != nil {
		log.Error(fmt.Errorf("cannot create server: %v", err))
		return
	}

	err = server.Serve()
	if err != nil {
		log.Error(fmt.Errorf("error during servant: %v", err))
	}
}
