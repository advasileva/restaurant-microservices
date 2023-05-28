package session

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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
		(*session)(nil),
	}

	for _, model := range models {
		err := r.db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return fmt.Errorf("cannot setup session table: %v", err)
		}
	}

	return nil
}
