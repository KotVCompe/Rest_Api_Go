package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"size:100;unique;not null" json:"email"`
	Age       int            `gorm:"not null" json:"age"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=0,max=150"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=2,max=100"`
	Email string `json:"email" binding:"omitempty,email"`
	Age   int    `json:"age" binding:"omitempty,min=0,max=150"`
}
