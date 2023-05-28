package user

import "time"

type user struct {
	Id           int64     // Id is automatically detected as primary key
	Username     string    `pg:"unique:username,notnull"`
	Email        string    `pg:"unique:email,notnull"`
	PasswordHash string    `pg:"notnull:password_hash"`
	Role         string    `pg:"notnull:role"`
	CreatedAt    time.Time `pg:"default:now()"`
	UpdatedAt    time.Time `pg:"default:now(),on_update:now()"`
}
