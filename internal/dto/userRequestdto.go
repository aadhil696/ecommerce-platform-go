package dto

import "go-ecommerce-app/internal/domain"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUp struct {
	UserLogin
	PhoneNo string `json:"phoneno"`
}

type UserVerifyCode struct {
	Code int `json:"code"`
}

type SellerInput struct {
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	PhoneNumber       string `json:"phoneno"`
	BankAccountNumber uint   `json:"bankaccountno"`
	SwiftCode         string `json:"swiftcode"`
	PaymentType       string `json:"paymenttype"`
}

type AddressInput struct {
	AddressLine1 string `json:"addressline1"`
	AddressLine2 string `json:"addressline2"`
	City         string `json:"city"`
	PostCode     uint   `json:"postcode"`
	Country      string `json:"country"`
}

type ProfileInput struct {
	FirstName    string       `json:"firstname"`
	LastName     string       `json:"lastname"`
	AddressInput AddressInput `json:"address"`
}

type UserProfileResponse struct {
	FirstName string         `json:"firstname"`
	LastName  string         `json:"lastname"`
	Email     string         `json:"email" gorm:"index;unique;not null"`
	Phone     string         `json:"phone"`
	Address   domain.Address `json:"address"` //relation
	Verified  bool           `json:"verified" gorm:"default:false"`
	UserType  string         `json:"usertype" gorm:"default:buyer"`
	Cart      domain.Cart    `json:"cart"`
	Orders    []domain.Order `json:"orders"`
}
