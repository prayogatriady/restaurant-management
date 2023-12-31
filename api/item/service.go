package item

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/prayogatriady/restaurant-management/model"
	"github.com/prayogatriady/restaurant-management/model/item_model"
)

type ItemService interface {
	ItemList(ctx context.Context) (response *model.BaseResponse, data []*item_model.ItemListResponse)

	GenDummyCategories(ctx context.Context) (response *model.BaseResponse, data *item_model.GenDummyCategoriesResponse)
	GenDummyItems(ctx context.Context, request *item_model.GenDummyItemsRequest) (response *model.BaseResponse, data *item_model.GenDummyItemsResponse)
}

type itemService struct {
	itemRepository ItemRepository
}

func NewItemService(repo ItemRepository) ItemService {
	return &itemService{
		itemRepository: repo,
	}
}

func (s *itemService) ItemList(ctx context.Context) (response *model.BaseResponse, data []*item_model.ItemListResponse) {

	var statusCode int = http.StatusOK
	var errors interface{}

	items, err := s.itemRepository.ItemList(ctx)
	if err != nil {
		statusCode = http.StatusInternalServerError
		errors = err.Error()
	}

	// Initialize a map to store items by category
	categoryMap := make(map[string]*item_model.ItemListResponse)

	// Iterate through the items and categorize them
	for _, item := range items {
		category, exists := categoryMap[item.CategoryName]

		if !exists {
			category = &item_model.ItemListResponse{Category: item.CategoryName}
		}

		category.Items = append(category.Items, item)
		categoryMap[item.CategoryName] = category
	}

	// Convert the map of categories to a slice
	for _, category := range categoryMap {
		data = append(data, category)
	}

	response = &model.BaseResponse{
		Status: statusCode,
		Errors: errors,
	}

	return response, data
}

var categoriesDummy = []string{"Ramen", "Sushi", "Udon", "Beverage"}

func (s *itemService) GenDummyCategories(ctx context.Context) (response *model.BaseResponse, data *item_model.GenDummyCategoriesResponse) {

	var categories []*item_model.Category
	var statusCode int = http.StatusCreated
	var message string = "Successfully added"
	var errors interface{}

	for _, category := range categoriesDummy {
		categories = append(categories, &item_model.Category{
			Name:        category,
			Description: fmt.Sprintf("Description of %s", category),
		})
	}

	err := s.itemRepository.GenDummyCategories(ctx, categories)
	if err != nil {
		statusCode = http.StatusInternalServerError
		message = ""
		errors = err.Error()
	}

	response = &model.BaseResponse{
		Status: statusCode,
		Errors: errors,
	}

	data = &item_model.GenDummyCategoriesResponse{
		Message: message,
	}

	return response, data
}

func (s *itemService) GenDummyItems(ctx context.Context, request *item_model.GenDummyItemsRequest) (response *model.BaseResponse, data *item_model.GenDummyItemsResponse) {

	var items []*item_model.Item
	var statusCode int = http.StatusCreated
	var message string = "Successfully added"
	var errors interface{}

	pricesDummy := []int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000}

	for catIdx, category := range categoriesDummy {
		for i := 1; i <= request.ItemAmount; i++ {

			randomIndex := rand.Intn(len(pricesDummy))

			items = append(items, &item_model.Item{
				Name:        fmt.Sprintf("%s %d", category, i),
				Description: fmt.Sprintf("Description of %s %d", category, i),
				Price:       pricesDummy[randomIndex],
				Quantity:    rand.Intn(20) + 21, // random integer between 21 and 40
				CategoryId:  catIdx + 1,
				IsActive:    true,
			})
		}
	}

	err := s.itemRepository.GenDummyItems(ctx, items)
	if err != nil {
		statusCode = http.StatusInternalServerError
		message = ""
		errors = err.Error()
	}

	response = &model.BaseResponse{
		Status: statusCode,
		Errors: errors,
	}

	data = &item_model.GenDummyItemsResponse{
		Message: message,
	}

	return response, data
}
