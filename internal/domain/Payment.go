package domain

import "time"

type Payment struct {
	ID            uint      `json:"id" gorm:"PrimaryKey"`
	UserId        uint      `json:"userid"`
	CaptureMethod string    `json:"capturemethod"`
	Amount        float64   `json:"amount"`
	TransactionId uint      `json:"transactionid"`
	CustomerId    string    `json:"customerid"`
	PaymentId     string    `json:"paymentid"`
	Status        string    `json:"status"`
	Response      string    `json:"response"`
	CreatedAt     time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"default:current_timestamp"`
}
