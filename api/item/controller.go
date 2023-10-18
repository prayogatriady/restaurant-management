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

	GenDummyCategories(c *gin.Context)
	GenDummyItems(c *gin.Context)
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

func (ic *itemController) GenDummyCategories(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, data := ic.itemService.GenDummyCategories(ctx)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "GenDummyCategories",
		Errors:      response.Errors,
	})
}

func (ic *itemController) GenDummyItems(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var request item_model.GenDummyItemsRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:  c,
			Data:        nil,
			StatusCode:  http.StatusBadRequest,
			ServiceName: "GenDummyItems",
			Errors:      "Bad Request",
		})
		return
	}

	response, data := ic.itemService.GenDummyItems(ctx, &request)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "GenDummyItems",
		Errors:      response.Errors,
	})
}
