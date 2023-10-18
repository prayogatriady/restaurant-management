package report_model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	Id          int
	ReportType  string
	StartDate   time.Time
	EndDate     time.Time
	TotalIncome int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type GetBillsReportParams struct {
	StartDate string
	EndDate   string
}

type GetBillsReport struct {
	TrxYear    int
	TrxMonth   int
	TotalPrice int
}
