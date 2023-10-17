package order

import (
	"context"

	"github.com/prayogatriady/restaurant-management/model/item_model"
	"github.com/prayogatriady/restaurant-management/model/order_model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	AddOrder(ctx context.Context, order *order_model.Order) (err error)
	AddOrderDetail(ctx context.Context, orders []*order_model.OrderDetail) (err error)
	UpdateItem(ctx context.Context, item *item_model.Item) (err error)
}

type orderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		Db: db,
	}
}

func (r *orderRepository) AddOrder(ctx context.Context, order *order_model.Order) (err error) {
	err = r.Db.WithContext(ctx).Table("t_order").Create(&order).Error
	return
}

func (r *orderRepository) AddOrderDetail(ctx context.Context, orders []*order_model.OrderDetail) (err error) {
	err = r.Db.WithContext(ctx).Table("t_order_detail").Create(&orders).Error
	return
}

func (r *orderRepository) UpdateItem(ctx context.Context, item *item_model.Item) (err error) {

	err = r.Db.WithContext(ctx).Table("m_item").Where("id = ?", item.Id).Update("quantity", gorm.Expr("quantity - ?", item.Quantity)).Error
	return

}
