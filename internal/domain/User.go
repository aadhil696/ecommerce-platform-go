package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"usertype" gorm:"default:buyer"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
