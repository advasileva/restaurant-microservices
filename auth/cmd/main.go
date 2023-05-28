package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"server/internal/database"
	"server/internal/http"
	"server/internal/repositories/user"
)

func main() {
	db := database.Connect()
	defer database.Disconnect(db)

	userRepository := user.NewRepository(db)

	server, err := http.NewServer(userRepository)
	if err != nil {
		log.Error(fmt.Errorf("cannot create server: %v", err))
	}

	err = server.Serve()
	if err != nil {
		log.Error(fmt.Errorf("error during servant: %v", err))
	}
}
