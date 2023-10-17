package item

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/restaurant-management/http/httpresponse"
	"github.com/prayogatriady/restaurant-management/model/item_model"
)

type ItemController interface {
	ItemList(c *gin.Context)

	AddBulkCategories(c *gin.Context)
	AddBulkItems(c *gin.Context)
}

type itemController struct {
	itemService ItemService
}

func NewItemController(itemService ItemService) ItemController {
	return &itemController{
		itemService: itemService,
	}
}

func (ic *itemController) ItemList(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, data := ic.itemService.ItemList(ctx)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "ItemList",
		Errors:      response.Errors,
	})
}

func (ic *itemController) AddBulkCategories(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, data := ic.itemService.AddBulkCategories(ctx)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "AddBulkCategories",
		Errors:      response.Errors,
	})
}

func (ic *itemController) AddBulkItems(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var request item_model.AddBulkItemsRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:  c,
			Data:        nil,
			StatusCode:  http.StatusBadRequest,
			ServiceName: "AddBulkItems",
			Errors:      "Bad Request",
		})
		return
	}

	response, data := ic.itemService.AddBulkItems(ctx, &request)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "AddBulkItems",
		Errors:      response.Errors,
	})
}
