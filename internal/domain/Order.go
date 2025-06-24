package domain

import "time"

type Order struct {
	ID             uint        `json:"id" gorm:"PrimaryKey"`
	UserId         uint        `json:"userid"`
	Status         string      `json:"status"`
	Amount         float64     `json:"amount"`
	PaymentId      string      `json:"paymentid"`
	TransactionId  string      `json:"transactionid"`
	OrderRefNumber int         `json:"orderrefnumber"`
	Items          []OrderItem `json:"items"`
	CreatedAt      time.Time   `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time   `gorm:"default:current_timestamp"`
}
