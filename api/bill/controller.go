package bill

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/restaurant-management/http/httpresponse"
	"github.com/prayogatriady/restaurant-management/model/bill_model"
)

type BillController interface {
	CreateBill(c *gin.Context)

	// GenDummyCategories(c *gin.Context)
}

type billController struct {
	billService BillService
}

func NewBillController(billService BillService) BillController {
	return &billController{
		billService: billService,
	}
}

func (ic *billController) CreateBill(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var request bill_model.CreateBillRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:  c,
			Data:        nil,
			StatusCode:  http.StatusBadRequest,
			ServiceName: "CreateBill",
			Errors:      "Bad Request",
		})
		return
	}

	response, data := ic.billService.CreateOrder(ctx, &request)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "CreateBill",
		Errors:      response.Errors,
	})
}
