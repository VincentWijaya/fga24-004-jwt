package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string    `gorm:"not null"`
	Email    string    `gorm:"not null"`
	Password string    `gorm:"not null"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateUserRequest struct {
	FullName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=7,lte=25"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=7,lte=25"`
}
