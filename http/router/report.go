package router

import "github.com/prayogatriady/restaurant-management/api/report"

func InjecttReport(routerCfg *RouterConfig) report.ReportController {
	reportRepository := report.NewReportRepository(routerCfg.Db)
	reportService := report.NewReportService(reportRepository)
	reportController := report.NewReportController(reportService)
	return reportController
}

func RouterReport(routerCfg *RouterConfig) {

	reportGroup := routerCfg.Api.Group("/report")

	reportGroup.POST("/create", InjecttReport(routerCfg).CreateReport)
}
