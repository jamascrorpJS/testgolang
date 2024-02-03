package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Yes        Status = "Да"
	No         Status = "Нет"
	InProccess Status = "В обработке"
)

type Operation struct {
	ID        uuid.UUID `gorm:"operation_id; primary_key" json:"id"`
	UserID    uuid.UUID `gorm:"user_id" json:"user_id"`
	User      User      `gorm:"foreign_key:user_id" json:"-"`
	Sum       float64   `gorm:"sum" json:"sum"`
	Comission int       `gorm:"perce" json:"comission"`
	Status    Status    `gorm:"status" json:"status"`
	gorm.Model
}

type TotalOperations struct {
	Total  int `gorm:"total" json:"total"`
	Amount int `gorm:"amount" json:"amount"`
}
