package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterConfig struct {
	Api *gin.RouterGroup
	Db  *gorm.DB
}

func ApiRoutes(routerCfg *RouterConfig) {

	api := routerCfg.Api.Group("/api")

	routerCfg.Api = api

	// RouterPing(routerCfg)
	RouterItem(routerCfg)
	RouterOrder(routerCfg)
}
