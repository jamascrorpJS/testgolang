package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jamascrorpJS/eBank/internal/domain"
	"github.com/jamascrorpJS/eBank/internal/repository"
)

type UserServiceI interface {
	Existed(ctx context.Context, phone string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}
type BalanceServiceI interface {
	CreateOperation(ctx context.Context, sum float64, comission int, id string) (uuid.UUID, error)
	UpdateBalance(ctx context.Context, id string, sum float64, comission int) error
	GetBalance(ctx context.Context, id string) (*float64, error)
}
type OperationServiceI interface {
	GetAllOperations(ctx context.Context) (domain.TotalOperations, error)
}
type Service struct {
	User      UserServiceI
	Balance   BalanceServiceI
	Operation OperationServiceI
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:      NewUserService(repository),
		Balance:   NewBalanceService(repository),
		Operation: NewOperationService(repository),
	}
}
