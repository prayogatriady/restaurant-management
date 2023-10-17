package item_model

type AddBulkItemsResponse struct {
	Message string `json:"message"`
}

type AddBulkCategoriesResponse struct {
	Message string `json:"message"`
}

type ItemListResponse struct {
	Category string      `json:"category"`
	Items    []*ItemList `json:"items"`
}
