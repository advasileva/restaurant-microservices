package order

import (
	"server/internal/models"
	"time"
)

type dish struct {
	Id          int64
	Name        string  `pg:"unique:name"`
	Description string  `pg:""`
	Price       float64 `pg:""`
	Quantity    int64   `pg:""`
}

func (d *dish) toModel() models.Dish {
	return models.Dish{
		Name:        d.Name,
		Description: d.Description,
		Price:       d.Price,
		Quantity:    d.Quantity,
	}
}

type order struct {
	Id              int64
	UserId          int64     `pg:""`
	Status          string    `pg:""`
	SpecialRequests string    `pg:""`
	CreatedAt       time.Time `pg:"default:now()"`
	UpdatedAt       time.Time `pg:"default:now(),on_update:now()"`
}

type orderDish struct {
	Id       int64
	OrderId  int64   `pg:""`
	DishId   int64   `pg:""`
	Quantity int64   `pg:""`
	Price    float64 `pg:""`
}
