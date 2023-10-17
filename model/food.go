package model

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	FoodId      int
	Name        string
	Description string
	Price       string
	Category    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
