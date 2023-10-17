package model

type BaseResponse struct {
	Status int         `json:"status"`
	Errors interface{} `json:"errors"`
}
