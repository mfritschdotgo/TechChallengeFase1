package dto

type ProductItem struct {
	ID       string `json:"id"`
	Quantity int64  `json:"quantity"`
}

type CreateOrderRequest struct {
	Client   string        `json:"client"`
	Products []ProductItem `json:"products"`
	Status   int           `json:"status"`
}
