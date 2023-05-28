package models

import "time"

type Order struct {
	UserId          int64     `json:"user_id"`
	Status          string    `json:"status"`
	SpecialRequests string    `json:"special_requests"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
