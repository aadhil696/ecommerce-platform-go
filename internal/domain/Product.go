package domain

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"index;"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"categoryid"`
	ImageUrl    string    `json:"imageurl"`
	Price       float64   `json:"price"`
	UserId      int       `json:"userid"`
	Stock       uint      `json:"stock"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"default:current_timestamp"`
}
