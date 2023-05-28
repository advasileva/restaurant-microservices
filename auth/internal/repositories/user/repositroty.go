package user

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"golang.org/x/crypto/bcrypt"
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
		(*user)(nil),
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

func (r *repository) Add(username string, email string, password string, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("cannot generate password hash: %v", err)
	}

	dto := &user{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
		Role:         role,
	}
	_, err = r.db.Model(dto).Insert()
	if err != nil {
		return fmt.Errorf("cannot insert row: %v", err)
	}

	return nil
}

func (r *repository) CheckPasswordByEmail(email string, password string) (bool, error) {
	var dto user
	err := r.db.Model(&dto).Column("password_hash").Where("email = ?", email).Select()
	if err != nil {
		return false, fmt.Errorf("cannot find row: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dto.PasswordHash), []byte(password))

	return err == nil, err
}

func (r *repository) GetUsernameByEmail(email string) (string, error) {
	var dto user
	err := r.db.Model(&dto).Column("username").Where("email = ?", email).Select()
	if err != nil {
		return "", fmt.Errorf("cannot find username by email: %v", err)
	}

	return dto.Username, nil
}
