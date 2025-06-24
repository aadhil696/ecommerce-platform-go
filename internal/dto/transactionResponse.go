package dto

type SellerOrderDetails struct {
	OrderrRefNumber int    `json:"orderrefnumber"`
	OrderStatus     int    `json:"order_status"`
	CreatedAt       string `json:"createdat"`
	OrderItemId     uint   `json:"orderitemid"`
	ProductId       uint   `json:"product_id"`
	Name            string `json:"name"`
	ImageUrl        string `json:"imageurl"`
	Price           string `json:"price"`
	Qty             uint   `json:"qty"`
	CustomerName    string `json:"customername"`
	CustomerEmail   string `json:"customeremail"`
	CustomerPhone   string `json:"customerphone"`
	CustomerAddress string `json:"customeraddress"`
}
