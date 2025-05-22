package domain

import "time"

type BankAccount struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserId      uint      `json:"userid"`
	BankAccount uint      `json:"bankaccount" gorm:"index;unique;not null"`
	SwiftCode   string    `json:"swiftcode"`
	PaymentType string    `json:"paymenttype"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"default:current_timestamp"`
}
