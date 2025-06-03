package dto

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageurl"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"categoryid"`
	Stock       uint    `json:"stock"`
}

type UpdateStockRequest struct {
	Stock uint `json:"stock"`
}
