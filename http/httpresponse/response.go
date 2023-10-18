package httpresponse

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/restaurant-management/utils/config"
)

type HttpParams struct {
	GinContext  *gin.Context
	Data        interface{}
	Payload     interface{}
	StatusCode  int
	ServiceName string
	Errors      interface{}
}

type Response struct {
	Code        int         `json:"code"`
	Data        interface{} `json:"data"`
	Errors      interface{} `json:"errors"`
	AppName     string      `json:"app_name"`
	AppVersion  string      `json:"app_version"`
	CurrentTime string      `json:"current_time"`
}

func BaseResponse(httpParams *HttpParams) {

	appConfig := config.AppCfg

	loc, _ := time.LoadLocation(appConfig.App.Timezone)
	currentTime := time.Now().In(loc)
	currentTimeNew := currentTime.Format(time.RFC3339)

	payload, err := json.Marshal(httpParams.Payload)
	if err != nil {
		return
	}
	data, err := json.Marshal(httpParams.Data)
	if err != nil {
		return
	}

	log.Printf("%s | %s | %s | %d | %s | %s | %v \n",
		currentTimeNew,
		httpParams.ServiceName,
		httpParams.GinContext.Request.Host+httpParams.GinContext.Request.URL.Path,
		httpParams.StatusCode,
		payload,
		data,
		httpParams.Errors)

	response := Response{
		Code:        httpParams.StatusCode,
		Data:        httpParams.Data,
		Errors:      httpParams.Errors,
		AppName:     appConfig.App.AppName,
		AppVersion:  appConfig.App.AppVersion,
		CurrentTime: currentTimeNew,
	}

	httpParams.GinContext.JSON(httpParams.StatusCode, &response)
}
