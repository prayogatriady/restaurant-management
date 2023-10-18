package item_model

type GenDummyItemsResponse struct {
	Message string `json:"message"`
}

type GenDummyCategoriesResponse struct {
	Message string `json:"message"`
}

type ItemListResponse struct {
	Category string      `json:"category"`
	Items    []*ItemList `json:"items"`
}
