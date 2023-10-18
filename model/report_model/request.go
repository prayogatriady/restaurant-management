package report_model

type CreateReportRequest struct {
	ReportType     string `json:"report_type"`
	StartDateYear  int    `json:"start_date_year"`
	StartDateMonth int    `json:"start_date_month"`
	EndDateYear    int    `json:"end_date_year"`
	EndDateMonth   int    `json:"end_date_month"`
}
