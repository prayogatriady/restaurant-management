package router

import "github.com/prayogatriady/restaurant-management/api/order"

func InjecttOrder(routerCfg *RouterConfig) order.OrderController {
	orderRepository := order.NewOrderRepository(routerCfg.Db)
	orderService := order.NewOrderService(orderRepository)
	orderController := order.NewOrderController(orderService)
	return orderController
}

func RouterOrder(routerCfg *RouterConfig) {

	orderGroup := routerCfg.Api.Group("/order")

	orderGroup.POST("/create", InjecttOrder(routerCfg).CreateOrder)
}
