package dto

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageurl"`
	Description string  `json:"description"`
	CategoryID  int     `json:"categoryid"`
	Stock       int     `json:"stock"`
}

type UpdateStockRequest struct {
	Stock int `json:"stock"`
}
