package dto

type CreateCartRequest struct {
	ProductId uint `json:"productid"`
	Qty       uint `json:"qty"`
}
