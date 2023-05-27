package database

import (
	"github.com/MogLuiz/key-user-api/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}
