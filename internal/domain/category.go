package domain

import "time"

type Category struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"index;"`
	ParentId     uint      `json:"parentid"`
	ImageUrl     string    `json:"imageurl"`
	Products     []Product `json:"products"`
	DisplayOrder int       `json:"displayorder"`
	CreatedAt    time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"default:current_timestamp"`
}
