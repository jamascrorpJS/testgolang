package repository

import (
	"context"

	"github.com/jamascrorpJS/eBank/internal/domain"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) CreateUser(ctx context.Context, user domain.User) error {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Existed(ctx context.Context, phone string) (domain.User, error) {
	user := domain.User{}
	err := u.db.WithContext(ctx).Model(&user).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
