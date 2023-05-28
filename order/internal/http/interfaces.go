package http

import "server/internal/models"

type orderRepository interface {
	AddOrder(userId int64, dishes []models.Dish, Status string, SpecialRequests string) error
	ProcessOrders() error
	GetOrder(id int64) (models.Order, error)
}
