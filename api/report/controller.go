package report

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/restaurant-management/http/httpresponse"
	"github.com/prayogatriady/restaurant-management/model/report_model"
)

type ReportController interface {
	CreateReport(c *gin.Context)
}

type reportController struct {
	reportService ReportService
}

func NewReportController(reportService ReportService) ReportController {
	return &reportController{
		reportService: reportService,
	}
}

func (ic *reportController) CreateReport(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var request report_model.CreateReportRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		httpresponse.BaseResponse(&httpresponse.HttpParams{
			GinContext:  c,
			Data:        nil,
			StatusCode:  http.StatusBadRequest,
			ServiceName: "CreateReport",
			Errors:      "Bad Request",
		})
		return
	}

	response, data := ic.reportService.CreateReport(ctx, &request)
	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		Data:        data,
		StatusCode:  response.Status,
		ServiceName: "CreateReport",
		Errors:      response.Errors,
	})
}
