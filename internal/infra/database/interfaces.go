package database

import "github.com/MogLuiz/key-user-api/internal/entity"

type IUser interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type IProduct interface {
	Create(product *entity.Product) error
	// FindAll(page, limit int, sort string) ([]*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	// Delete(id string) error
}
