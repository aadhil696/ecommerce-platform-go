package domain

import "time"

type Cart struct {
	ID        int       `json:"id" gorm:"PrimaryKey"`
	UserId    int       `json:"userid"`
	ProductId int       `json:"productid"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"imageurl"`
	Price     float64   `json:"price"`
	Qty       int       `json:"qty"`
	SellerId  int       `json:"sellerid"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
