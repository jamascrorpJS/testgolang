package service

import (
	"context"
	"errors"

	"github.com/jamascrorpJS/eBank/internal/domain"
	"github.com/jamascrorpJS/eBank/internal/repository"
)

type UserService struct {
	repository repository.UserInterface
}

func NewUserService(repository *repository.Repository) *UserService {
	return &UserService{repository: repository.User}
}

func (u *UserService) Existed(ctx context.Context, phone string) (domain.User, error) {
	return u.repository.Existed(ctx, phone)
}

func (u *UserService) CreateUser(ctx context.Context, user domain.User) error {
	err := user.Validate()
	if err != nil {
		return errors.New("No valid user field")
	}
	return u.repository.CreateUser(ctx, user)
}
