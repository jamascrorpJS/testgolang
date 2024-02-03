package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sex string

const (
	Male       Sex = "Мужской"
	Female     Sex = "Женский"
	NotIdentif Sex = "Не идентифицирует"
)

type User struct {
	ID              uuid.UUID `gorm:"user_id; primary_key" json:"id"`
	CurrentBalance  float64   `gorm:"balance" json:"balance"`
	IsIdentificated bool      `gorm:"isIdentificated" json:"identificate"`
	Sex             Sex       `gorm:"sex" json:"sex" validate:"oneof=Мужской Женский Не идентифицирует"`
	Phone           string    `gorm:"phone; unique" json:"phone" validate:"numeric,len=9"`
	Name            string    `gorm:"name" json:"name"`
	LastName        string    `gorm:"lastname" json:"lastName"`
	FatherName      string    `gorm:"fathername" json:"fatherName"`
	gorm.Model
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
