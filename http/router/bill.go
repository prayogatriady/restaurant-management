package router

import "github.com/prayogatriady/restaurant-management/api/bill"

func InjecttBill(routerCfg *RouterConfig) bill.BillController {
	billRepository := bill.NewBillRepository(routerCfg.Db)
	billService := bill.NewBillService(billRepository)
	billController := bill.NewBillController(billService)
	return billController
}

func RouterBill(routerCfg *RouterConfig) {

	billGroup := routerCfg.Api.Group("/bill")

	billGroup.POST("/create", InjecttBill(routerCfg).CreateBill)
}
