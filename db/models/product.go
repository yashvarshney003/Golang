package models

import "time"

type product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProductID   string `gorm: "type:text"`
	VariantID   string `gorm: "type:text"`
}
