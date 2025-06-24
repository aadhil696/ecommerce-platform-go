package domain

import "time"

type OrderItem struct {
	ID        int       `json:"id" gorm:"PrimaryKey"`
	ProductId int       `json:"productid"`
	OrderId   int       `json:"orderid"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"imageurl"`
	Price     float64   `json:"price"`
	Qty       int       `json:"qty"`
	SellerId  int       `json:"sellerid"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:current_timestamp"`
}
