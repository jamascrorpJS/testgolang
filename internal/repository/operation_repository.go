package repository

import (
	"context"
	"time"

	"github.com/jamascrorpJS/eBank/internal/domain"
	"gorm.io/gorm"
)

type Operation struct {
	db *gorm.DB
}

func NewOperation(db *gorm.DB) *Operation {
	return &Operation{
		db: db,
	}
}

func (o *Operation) GetAllOperations(ctx context.Context, date time.Time) (domain.TotalOperations, error) {
	model := domain.TotalOperations{}
	err := o.db.WithContext(ctx).
		Table("operations").
		Select("COUNT(sum) as total, SUM(sum) as amount").
		Where("status = ? AND created_at >= ?", "Да", date).Find(&model).Error

	if err != nil {
		return model, err
	}
	return model, nil
}
