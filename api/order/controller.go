package order

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/restaurant-management/http/httpresponse"
	"github.com/prayogatriady/restaurant-management/model/order_model"
)

type OrderController interface {
	CreateOrder(c *gin.Context)
}

type orderController struct {
	orderService OrderService
}

func NewOrderController(orderService OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

func (oc *orderController) CreateOrder(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var request []*order_model.CreateOrderRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:  c,
			Data:        nil,
			StatusCode:  http.StatusBadRequest,
			ServiceName: "CreateOrder",
			Errors:      "Bad Request",
		})
		return
	}

	response, data := oc.orderService.CreateOrder(ctx, request)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "CreateOrder",
		Errors:      response.Errors,
	})
}
