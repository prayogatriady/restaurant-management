package bill

import (
	"context"

	"github.com/prayogatriady/restaurant-management/model/bill_model"
	"github.com/prayogatriady/restaurant-management/model/order_model"
	"gorm.io/gorm"
)

type BillRepository interface {
	CreateBill(ctx context.Context, bill *bill_model.Bill) (err error)
	GetOrder(ctx context.Context, orderId int) (order *order_model.Order, err error)

	GenDummyBills(ctx context.Context, bills []*bill_model.Bill) (err error)
	GenDummyOrders(ctx context.Context, orders []*order_model.Order) (err error)
}

type billRepository struct {
	Db *gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{
		Db: db,
	}
}

func (r *billRepository) CreateBill(ctx context.Context, bill *bill_model.Bill) (err error) {
	err = r.Db.WithContext(ctx).Table("t_bill").Create(&bill).Error
	return
}

func (r *billRepository) GetOrder(ctx context.Context, orderId int) (order *order_model.Order, err error) {
	err = r.Db.WithContext(ctx).Table("t_order").Where("id = ?", orderId).Find(&order).Error
	return
}

func (r *billRepository) GenDummyBills(ctx context.Context, bills []*bill_model.Bill) (err error) {
	err = r.Db.WithContext(ctx).Table("t_bill").Create(&bills).Error
	return
}

func (r *billRepository) GenDummyOrders(ctx context.Context, orders []*order_model.Order) (err error) {
	err = r.Db.WithContext(ctx).Table("t_order").Create(&orders).Error
	return
}
