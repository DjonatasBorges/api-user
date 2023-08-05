package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" validate:"required,max=100"`
	Lastname string    `json:"lastname" validate:"required,max=100"`
	Cpf      string    `json:"cpf" validate:"required,len=11"`
	Email    string    `json:"email" validate:"required,email"`
	Phone    string    `json:"phone" validate:"required,max=20"`
	Password string    `json:"password" validate:"required,min=6,max=100"`
}
