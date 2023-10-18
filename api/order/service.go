package order

import (
	"context"
	"net/http"
	"time"

	"github.com/prayogatriady/restaurant-management/model"
	"github.com/prayogatriady/restaurant-management/model/item_model"
	"github.com/prayogatriady/restaurant-management/model/order_model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request []*order_model.CreateOrderRequest) (response *model.BaseResponse, data *order_model.CreateOrderResponse)
}

type orderService struct {
	orderRepository OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{
		orderRepository: repo,
	}
}

func (s *orderService) CreateOrder(
	ctx context.Context,
	request []*order_model.CreateOrderRequest) (response *model.BaseResponse, data *order_model.CreateOrderResponse) {

	var totalPrice int
	for _, r := range request {
		totalPrice += r.Price
	}

	order := &order_model.Order{
		TotalPrice:    totalPrice,
		OrderDatetime: time.Now(),
	}

	if err := s.orderRepository.AddOrder(ctx, order); err != nil {
		response = &model.BaseResponse{
			Status: http.StatusInternalServerError,
			Errors: err.Error(),
		}
		return
	}

	var orderDetail []*order_model.OrderDetail
	for _, r := range request {

		item := &item_model.Item{
			Id:       r.ItemId,
			Quantity: r.Quantity,
		}
		s.orderRepository.UpdateItem(ctx, item)

		orderDetail = append(orderDetail, &order_model.OrderDetail{
			OrderId:  order.Id,
			ItemId:   r.ItemId,
			Quantity: r.Quantity,
			Price:    r.Price,
		})
	}

	if err := s.orderRepository.AddOrderDetail(ctx, orderDetail); err != nil {
		response = &model.BaseResponse{
			Status: http.StatusInternalServerError,
			Errors: err.Error(),
		}
		return
	}

	response = &model.BaseResponse{
		Status: http.StatusCreated,
	}

	data = &order_model.CreateOrderResponse{
		Message: "Successfully created order",
	}

	return
}
