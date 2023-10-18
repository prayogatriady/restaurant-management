package bill_model

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	Id           int
	OrderId      int
	TotalPrice   int
	BillDatetime time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
