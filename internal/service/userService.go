package service

import (
	"errors"
	"fmt"
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/pkg/notification"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config configs.AppConfig
}

// func (s UserService) findUserByEmail(email string) (*domain.User, error) {

// 	return &domain.User{}, nil
// }

func (s *UserService) SignUp(input *dto.UserSignUp) (string, error) {

	hPassword, err := s.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		return "", err
	}
	//Convert DTO to domain model
	// user := domain.User{
	// 	Email:    input.Email,
	// 	Password: hPassword,
	// 	Phone:    input.PhoneNo,
	// }

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.PhoneNo,
	})

	if err != nil {
		return "", fmt.Errorf("signup failed :%w", err)
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)

}

func (s *UserService) Login(input *dto.UserLogin) (string, error) {
	var user domain.User

	user, err := s.Repo.FindUser(input.Email)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	err = s.Auth.VerifyPassword(input.Password, user.Password)

	if err != nil {
		return "", errors.New("incorrect password")
	}

	//generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s *UserService) isVerified(id uint) bool {

	currentUser, err := s.Repo.FindUserbyID(id)

	return err == nil && currentUser.Verified
}

func (s *UserService) GetVerificationCode(e domain.User) (int, error) {

	//if user already verified
	if s.isVerified(e.ID) {
		log.Info("1st")
		return 0, errors.New("user already verified")
	}

	//generate verification code
	code, err := s.Auth.GenerateCode()
	if err != nil {
		return 0, fmt.Errorf("verification code generating failed due to %s", err)
	}

	//update user
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return 0, fmt.Errorf("user updation failed during generating verification code due to %s", err)
	}

	user, err = s.Repo.FindUserbyID(e.ID)
	if err != nil {
		return 0, err
	}
	//Send SMS
	notificationClient := notification.NewNotificationClient(s.Config)
	notificationClient.SendVoiceCall(user.Phone, strconv.Itoa(code))

	//return verification code

	return code, nil
}

func (s *UserService) VerifyCode(id uint, code int) error {
	if s.isVerified(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.FindUserbyID(id)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Code != code {
		return errors.New("verification code mismatch")
	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(user.ID, updateUser)
	if err != nil {
		return errors.New("verifiying code failed")
	}

	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (s *UserService) GetProfile(id uint) (*domain.User, error) {

	return &domain.User{}, nil
}

func (s *UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s *UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {
	//find exisiting user
	user, _ := s.Repo.FindUserbyID(id)

	//return already joined seller program
	if user.UserType == domain.SELLER {
		return "", errors.New("user is already upgraded as seller")
	}

	//update user
	seller, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	})
	if err != nil {
		return "", err
	}

	//generating token
	token, err := s.Auth.GenerateToken(user.ID, user.Email, seller.UserType)
	if err != nil {
		return "", fmt.Errorf("seller token generation failed %s", err)
	}

	//create bank account information
	err = s.Repo.AddBankAccount(domain.BankAccount{
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymentType,
		UserId:      id,
	})

	return token, err
}

func (s *UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (s *UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s *UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s *UserService) GetOrders(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s *UserService) GetOrderById(id uint) (*domain.User, error) {

	return &domain.User{}, nil
}
