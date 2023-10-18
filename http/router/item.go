package router

import "github.com/prayogatriady/restaurant-management/api/item"

func InjecttItem(routerCfg *RouterConfig) item.ItemController {
	itemRepository := item.NewItemRepository(routerCfg.Db)
	itemService := item.NewItemService(itemRepository)
	itemController := item.NewItemController(itemService)
	return itemController
}

func RouterItem(routerCfg *RouterConfig) {

	itemGroup := routerCfg.Api.Group("/item")

	itemGroup.GET("/itemList", InjecttItem(routerCfg).ItemList)
	itemGroup.POST("/genDummyCategories", InjecttItem(routerCfg).GenDummyCategories)
	itemGroup.POST("/genDummyItems", InjecttItem(routerCfg).GenDummyItems)
}
