package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jamascrorpJS/eBank/internal/domain"
	"github.com/jamascrorpJS/eBank/internal/repository"
)

const (
	MaxBalanceNoIdentification = 10000
	MaxBalanceIdentification   = 100000
)

type BalanceService struct {
	repository repository.BalanceInterface
}

func NewBalanceService(repository *repository.Repository) *BalanceService {
	return &BalanceService{repository: repository.Balance}
}

func (b *BalanceService) CreateOperation(ctx context.Context, sum float64, comission int, userID string) (uuid.UUID, error) {
	id, err := b.repository.CreateOperation(ctx, sum, comission, userID)
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

func (b *BalanceService) UpdateBalance(ctx context.Context, userID string, sum float64, comission int) error {

	currentBalance, err := b.GetBalance(ctx, userID)
	updatedBalance := (sum - (sum * float64(comission) / 100)) + *currentBalance

	id, err := b.CreateOperation(ctx, sum, comission, userID)
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Second)

	statusIdent, err := b.repository.Identificate(userID)
	if err != nil {
		return err
	}

	ok := validators(statusIdent, updatedBalance)

	if !ok {
		b.repository.UpdateBalance(ctx, id, updatedBalance, string(domain.No), userID, ok)
		return errors.New("Limit overflow")
	}

	err = b.repository.UpdateBalance(ctx, id, updatedBalance, string(domain.Yes), userID, ok)
	if err != nil {
		b.repository.UpdateBalance(ctx, id, updatedBalance, string(domain.No), userID, ok)
		return err
	}

	return nil
}

func (bs *BalanceService) GetBalance(ctx context.Context, userID string) (*float64, error) {
	return bs.repository.GetBalance(ctx, userID)
}

func validators(identificates bool, balance float64) bool {
	if identificates && balance > MaxBalanceIdentification {
		return false
	}
	if !identificates && balance > MaxBalanceNoIdentification {
		return false
	}
	return true
}
