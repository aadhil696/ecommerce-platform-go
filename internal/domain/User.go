package domain

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        int       `json:"id" gorm:"PrimaryKey"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Address   Address   `json:"address"` //relation
	Cart      Cart      `json:"cart"`    //relation
	Orders    []Order   `json:"orders"`  //relation
	Payments  []Payment `json:"payment"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"usertype" gorm:"default:buyer"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:current_timestamp"`
}
