package bill

import (
	"context"
	"net/http"
	"time"

	"github.com/prayogatriady/restaurant-management/model"
	"github.com/prayogatriady/restaurant-management/model/bill_model"
)

type BillService interface {
	CreateOrder(ctx context.Context, request *bill_model.CreateBillRequest) (response *model.BaseResponse, data *bill_model.CreateBillResponse)

	// GenDummyBills(ctx context.Context) (response *model.BaseResponse, data *bill_model.GenDummyCategoriesResponse)
}

type billService struct {
	billRepository BillRepository
}

func NewBillService(repo BillRepository) BillService {
	return &billService{
		billRepository: repo,
	}
}

func (s *billService) CreateOrder(ctx context.Context, request *bill_model.CreateBillRequest) (response *model.BaseResponse, data *bill_model.CreateBillResponse) {

	order, err := s.billRepository.GetOrder(ctx, request.OrderId)
	if err != nil {
		response = &model.BaseResponse{
			Status: http.StatusBadRequest,
			Errors: err.Error(),
		}
		return
	}

	bill := &bill_model.Bill{
		OrderId:      order.Id,
		TotalPrice:   order.TotalPrice,
		BillDatetime: time.Now(),
	}

	if err := s.billRepository.CreateBill(ctx, bill); err != nil {
		response = &model.BaseResponse{
			Status: http.StatusInternalServerError,
			Errors: err.Error(),
		}
		return
	}

	response = &model.BaseResponse{
		Status: http.StatusCreated,
	}

	data = &bill_model.CreateBillResponse{
		Message: "Successfully created bill",
	}

	return

}
