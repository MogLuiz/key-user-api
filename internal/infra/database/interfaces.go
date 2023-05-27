package database

import "github.com/MogLuiz/key-user-api/internal/entity"

type IUser interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
