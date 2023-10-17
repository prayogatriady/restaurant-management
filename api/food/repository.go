package food

import (
	"github.com/prayogatriady/restaurant-management/model"
)

type FoodRepository interface {
	GetFoodList(pagination *model.Pagination) (contacts []*model.Food, totalRows int64, err error)
	GetFood(contactId int) (contact *model.Food, err error)
	AddFood(contact *model.Food) (err error)
	EditFood(contact *model.Food) (err error)
	DeleteFood(contactId int) (err error)
}

type foodRepository struct{}
