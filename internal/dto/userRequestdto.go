package dto

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

type ProfileInput struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	AddressInput string `json:"addressinput"`
}
