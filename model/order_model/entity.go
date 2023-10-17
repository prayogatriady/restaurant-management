package order_model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id            int
	TotalPrice    int
	OrderDatetime time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type OrderDetail struct {
	Id        int
	OrderId   int
	ItemId    int
	Quantity  int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
