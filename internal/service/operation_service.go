package service

import (
	"context"
	"time"

	"github.com/jamascrorpJS/eBank/internal/domain"
	"github.com/jamascrorpJS/eBank/internal/repository"
)

type OperationService struct {
	repository repository.OperationInterface
}

func NewOperationService(repository *repository.Repository) *OperationService {
	return &OperationService{repository: repository.Operation}
}

func (os *OperationService) GetAllOperations(ctx context.Context) (domain.TotalOperations, error) {
	year, month, _ := time.Now().Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return os.repository.GetAllOperations(ctx, firstOfMonth)
}
