package main

import (
	"fmt"
	"server/internal/database"
	"server/internal/http"
	"server/internal/repositories/session"
	"server/internal/repositories/user"
	"server/internal/services"

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

	userRepository := user.NewRepository(db)
	err := userRepository.SetupTable()
	if err != nil {
		log.Error(fmt.Errorf("cannot setup table in user repository: %v", err))
		return
	}

	sessionRepository := session.NewRepository(db)
	err = sessionRepository.SetupTable()
	if err != nil {
		log.Error(fmt.Errorf("cannot setup table in session repository: %v", err))
		return
	}

	authService := services.NewAuthService()

	server, err := http.NewServer(userRepository, authService)
	if err != nil {
		log.Error(fmt.Errorf("cannot create server: %v", err))
		return
	}

	err = server.Serve()
	if err != nil {
		log.Error(fmt.Errorf("error during servant: %v", err))
	}
}
