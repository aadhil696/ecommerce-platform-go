package domain

import "time"

type Address struct {
	ID           int       `json:"id" gorm:"PrimaryKey"`
	AddressLine1 string    `json:"addressline1"`
	AddressLine2 string    `json:"addressline2"`
	City         string    `json:"city"`
	PostCode     uint      `json:"postcode"`
	Country      string    `json:"country"`
	UserID       int       `json:"userid"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}
