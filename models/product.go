package models

import "time"

type Products struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      Language   `json:"name"`
	Price     float64    `json:"price" validate:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Language struct {
	En string `json:"en"`
	Th string `json:"th"`
}