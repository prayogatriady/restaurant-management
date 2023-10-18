package report

import (
	"context"
	"net/http"
	"time"

	"github.com/prayogatriady/restaurant-management/model"
	"github.com/prayogatriady/restaurant-management/model/report_model"
	"github.com/prayogatriady/restaurant-management/utils/config"
)

type ReportService interface {
	CreateReport(ctx context.Context, request *report_model.CreateReportRequest) (response *model.BaseResponse, data []*report_model.CreateReportResponse)

	// GenDummyReports(ctx context.Context) (response *model.BaseResponse, data *report_model.GenDummyCategoriesResponse)
}

type reportService struct {
	reportRepository ReportRepository
}

func NewReportService(repo ReportRepository) ReportService {
	return &reportService{
		reportRepository: repo,
	}
}

func (s *reportService) CreateReport(
	ctx context.Context,
	request *report_model.CreateReportRequest) (response *model.BaseResponse, data []*report_model.CreateReportResponse) {

	var appCfg = config.AppCfg

	loc, _ := time.LoadLocation(appCfg.App.Timezone)

	startDate := time.Date(request.StartDateYear, time.Month(request.StartDateMonth), 1, 0, 0, 0, 0, loc).Format("2006-01-02")

	endDateNextMonth := time.Date(request.EndDateYear, time.Month(request.EndDateMonth+1), 1, 0, 0, 0, 0, loc)
	endDate := endDateNextMonth.AddDate(0, 0, -1).Format("2006-01-02")

	params := &report_model.GetBillsReportParams{
		StartDate: startDate,
		EndDate:   endDate,
	}

	bills, err := s.reportRepository.GetBills(ctx, params)
	if err != nil {
		response = &model.BaseResponse{
			Status: http.StatusBadRequest,
			Errors: err.Error(),
		}
		return
	}

	for _, bill := range bills {
		startDateBill := time.Date(bill.TrxYear, time.Month(bill.TrxMonth), 1, 0, 0, 0, 0, loc)
		endDateBill := startDateBill.AddDate(0, 1, -1)

		report := &report_model.Report{
			ReportType:  request.ReportType,
			StartDate:   startDateBill,
			EndDate:     endDateBill,
			TotalIncome: bill.TotalPrice,
		}

		if err := s.reportRepository.FindReport(ctx, report); err != nil {
			s.reportRepository.CreateReport(ctx, report)
		} else {
			s.reportRepository.UpdateReport(ctx, report)
		}

	}

	if err := s.reportRepository.UpdateReportedBill(ctx); err != nil {
		response = &model.BaseResponse{
			Status: http.StatusInternalServerError,
			Errors: err.Error(),
		}
		return
	}

	reports, err := s.reportRepository.FindReports(ctx)
	if err != nil {
		response = &model.BaseResponse{
			Status: http.StatusInternalServerError,
			Errors: err.Error(),
		}
		return
	}

	for _, report := range reports {
		data = append(data, &report_model.CreateReportResponse{
			ReportType:  report.ReportType,
			StartDate:   report.StartDate.Format("2006-01-02"),
			EndDate:     report.EndDate.Format("2006-01-02"),
			TotalIncome: report.TotalIncome,
		})
	}

	response = &model.BaseResponse{
		Status: http.StatusCreated,
	}

	return

}
