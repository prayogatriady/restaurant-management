package item_model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Id          int
	Name        string
	Description string
	Price       int
	CategoryId  int
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Category struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
