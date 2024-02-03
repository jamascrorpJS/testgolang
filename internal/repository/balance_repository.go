package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jamascrorpJS/eBank/internal/domain"
	"gorm.io/gorm"
)

type Balance struct {
	db *gorm.DB
}

func NewBalance(db *gorm.DB) *Balance {
	return &Balance{db: db}
}

func (b *Balance) CreateOperation(ctx context.Context, sum float64, comission int, userID string) (uuid.UUID, error) {
	operationID := uuid.New()
	id := uuid.MustParse(userID)
	operation := &domain.Operation{
		ID:        operationID,
		Sum:       sum,
		Comission: comission,
		Status:    domain.InProccess,
		UserID:    id,
	}

	err := b.db.WithContext(ctx).Create(operation).Error
	if err != nil {
		return uuid.UUID{}, err
	}
	return operation.ID, nil
}

func (b *Balance) UpdateBalance(
	ctx context.Context,
	operationID uuid.UUID,
	sum float64,
	status string,
	userID string,
	upd bool) error {
	tx := b.db.Begin()

	if !upd {
		rowsAffect := b.db.WithContext(ctx).Model(&domain.Operation{}).Where("id = ?", operationID).Update("status", status).RowsAffected
		if rowsAffect == 0 {
			return errors.New("No update")
		}
		return errors.New("No update")
	}

	rowsAffect := tx.WithContext(ctx).Model(&domain.Operation{}).Where("id = ?", operationID).Update("status", status).RowsAffected
	if rowsAffect == 0 {
		tx.Rollback()
		return errors.New("No update")
	}

	rowsAffect = tx.WithContext(ctx).Model(&domain.User{}).Where("id = ?", userID).Update("current_balance", sum).RowsAffected
	if rowsAffect == 0 {
		tx.Rollback()
		return errors.New("No update")
	}

	tx.Commit()
	return nil
}

func (b *Balance) GetBalance(ctx context.Context, id string) (*float64, error) {
	user := domain.User{}
	err := b.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user.CurrentBalance, nil
}

func (b *Balance) Identificate(id string) (bool, error) {
	user, err := b.getUser(id)
	if err != nil {
		return false, nil
	}
	return user.IsIdentificated, nil
}

func (b *Balance) getUser(id string) (domain.User, error) {
	user := domain.User{}
	err := b.db.Model(&user).Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
