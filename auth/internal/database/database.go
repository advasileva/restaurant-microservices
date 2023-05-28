package database

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

func Connect() *pg.DB {
	addr := fmt.Sprintf("%s:%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"))
	return pg.Connect(&pg.Options{
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Addr:     addr,
	})
}

func Disconnect(db *pg.DB) error {
	return db.Close()
}
