package report_model

type CreateReportResponse struct {
	ReportType  string `json:"report_type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	TotalIncome int    `json:"total_income"`
}
