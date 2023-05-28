package models

type Dish struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}
