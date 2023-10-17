package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Id          int64
	Name        string
	Description string
	Price       string
	Category    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
