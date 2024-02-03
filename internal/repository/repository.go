package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jamascrorpJS/eBank/internal/domain"
	"gorm.io/gorm"
)

type BalanceInterface interface {
	CreateOperation(ctx context.Context, sum float64, comission int, id string) (uuid.UUID, error)
	UpdateBalance(ctx context.Context, operationID uuid.UUID, sum float64, status string, userID string, upd bool) error
	GetBalance(ctx context.Context, id string) (*float64, error)
	Identificate(phone string) (bool, error)
}

type OperationInterface interface {
	GetAllOperations(ctx context.Context, date time.Time) (domain.TotalOperations, error)
}

type UserInterface interface {
	CreateUser(ctx context.Context, user domain.User) error
	Existed(ctx context.Context, phone string) (domain.User, error)
}

type Repository struct {
	Balance   BalanceInterface
	Operation OperationInterface
	User      UserInterface
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Balance:   NewBalance(db),
		Operation: NewOperation(db),
		User:      NewUser(db),
	}
}
