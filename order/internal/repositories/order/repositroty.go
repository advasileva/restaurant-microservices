package order

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/labstack/gommon/log"
	"server/internal/models"
)

func NewRepository(db *pg.DB) *repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *pg.DB
}

func (r *repository) SetupTable() error {
	models := []interface{}{
		(*dish)(nil),
		(*order)(nil),
		(*orderDish)(nil),
	}

	for _, model := range models {
		err := r.db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return fmt.Errorf("cannot setup users table: %v", err)
		}
	}

	return nil
}

func (r *repository) AddOrder(userId int64, dishes []models.Dish, status string, specialRequests string) error {
	dto := &order{
		UserId:          userId,
		Status:          status,
		SpecialRequests: specialRequests,
	}

	_, err := r.db.Model(dto).Insert()
	if err != nil {
		return fmt.Errorf("cannot insert row to order table: %v", err)
	}

	for _, dish := range dishes {
		err = r.addDish(dish.Name, dish.Description, dish.Price, dish.Quantity, dto.Id)
		if err != nil {
			return fmt.Errorf("cannot add dish relation: %v", err)
		}
	}

	return nil
}

func (r *repository) addDish(name string, description string, price float64, quantity int64, orderId int64) error {
	dto := &dish{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}

	_, err := r.db.Model(dto).Insert()
	if err != nil {
		log.Infof("dish with name=%s already exist", name)
	}

	err = r.addOrderDish(orderId, dto.Id, quantity, price)
	if err != nil {
		return fmt.Errorf("cannot add order_dish relation: %v", err)
	}

	return nil
}

func (r *repository) addOrderDish(orderId int64, dishId int64, quantity int64, price float64) error {
	dto := &orderDish{
		OrderId:  orderId,
		DishId:   dishId,
		Quantity: quantity,
		Price:    price,
	}

	_, err := r.db.Model(dto).Insert()
	if err != nil {
		return fmt.Errorf("cannot insert row to order_dish table: %v", err)
	}

	return nil
}

func (r *repository) ProcessOrders() error {
	var dto []order
	err := r.db.Model(&dto).Where("status = ?", "waiting").Select()
	if err != nil {
		return fmt.Errorf("cannot process orders: %v", err)
	}

	for _, order := range dto {
		order.Status = "ready"
		_, err = r.db.Model(&order).WherePK().Update()
		if err != nil {
			return fmt.Errorf("cannot update: %v", err)
		}
	}

	return nil
}

func (r *repository) GetOrder(id int64) (models.Order, error) {
	dto := &order{
		Id: id,
	}
	err := r.db.Model(dto).WherePK().Select()
	if err != nil {
		return models.Order{}, fmt.Errorf("cannot get order row: %v", err)
	}

	return models.Order{
		UserId:          dto.UserId,
		Status:          dto.Status,
		SpecialRequests: dto.SpecialRequests,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
	}, nil
}
