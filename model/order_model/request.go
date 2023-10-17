package order_model

type CreateOrderRequest struct {
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}
