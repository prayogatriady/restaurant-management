package report

import (
	"context"

	"github.com/prayogatriady/restaurant-management/model/report_model"
	"gorm.io/gorm"
)

type ReportRepository interface {
	CreateReport(ctx context.Context, report *report_model.Report) (err error)
	UpdateReport(ctx context.Context, report *report_model.Report) (err error)
	FindReport(ctx context.Context, report *report_model.Report) (err error)
	FindReports(ctx context.Context) (report []*report_model.Report, err error)
	GetBills(ctx context.Context, params *report_model.GetBillsReportParams) (bills []*report_model.GetBillsReport, err error)
}

type reportRepository struct {
	Db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{
		Db: db,
	}
}

func (r *reportRepository) CreateReport(ctx context.Context, report *report_model.Report) (err error) {
	err = r.Db.WithContext(ctx).Table("t_report").Create(&report).Error
	return
}

func (r *reportRepository) UpdateReport(ctx context.Context, report *report_model.Report) (err error) {

	startDate := report.StartDate.Format("2006-01-02")
	endDate := report.EndDate.Format("2006-01-02")

	err = r.Db.WithContext(ctx).Table("t_report").
		Where("start_date = ? AND end_date = ?", startDate, endDate).
		Update("total_income", gorm.Expr("total_income + ?", report.TotalIncome)).Error

	return
}

func (r *reportRepository) FindReport(ctx context.Context, report *report_model.Report) (err error) {

	startDate := report.StartDate.Format("2006-01-02")
	endDate := report.EndDate.Format("2006-01-02")

	err = r.Db.WithContext(ctx).Table("t_report").
		Where("start_date = ? AND end_date = ?", startDate, endDate).
		First(&report).Error

	return
}

func (r *reportRepository) FindReports(ctx context.Context) (report []*report_model.Report, err error) {
	err = r.Db.WithContext(ctx).Table("t_report").Find(&report).Error
	return
}

func (r *reportRepository) GetBills(ctx context.Context, params *report_model.GetBillsReportParams) (bills []*report_model.GetBillsReport, err error) {

	query := `
	SELECT
		YEAR(bill_datetime) AS trx_year,
		MONTH(bill_datetime) AS trx_month,
		SUM(total_price) AS total_price
	FROM t_bill
	WHERE bill_datetime BETWEEN ? AND ?
	GROUP BY YEAR(bill_datetime), MONTH(bill_datetime)
	ORDER BY trx_year, trx_month;
	`
	err = r.Db.WithContext(ctx).Raw(query, params.StartDate, params.EndDate).Scan(&bills).Error
	return

}
