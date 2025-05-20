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
