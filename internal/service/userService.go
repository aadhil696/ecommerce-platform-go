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
	"log"
	"strconv"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	CRepo  repository.CatalogRepository
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

func (s *UserService) isVerified(id int) bool {

	currentUser, err := s.Repo.FindUserbyID(id)

	return err == nil && currentUser.Verified
}

func (s *UserService) GetVerificationCode(e domain.User) (int, error) {

	//if user already verified
	if s.isVerified(e.ID) {
		log.Printf("1st")
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

func (s *UserService) VerifyCode(id int, code int) error {
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

func (s *UserService) CreateProfile(id int, input *dto.ProfileInput) error {

	///update user firstname and lastname
	_, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	})
	if err != nil {
		return err
	}

	//createprofile
	if err := s.Repo.CreateProfile(domain.Address{
		AddressLine1: input.AddressInput.AddressLine1,
		AddressLine2: input.AddressInput.AddressLine2,
		City:         input.AddressInput.City,
		PostCode:     input.AddressInput.PostCode,
		Country:      input.AddressInput.Country,
		UserID:       id,
	}); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetProfile(id uint) (*dto.UserProfileResponse, error) {

	profile, err := s.Repo.FindUserbyID(int(id))
	if err != nil {
		return nil, err
	}

	userProfile := dto.UserProfileResponse{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Phone:     profile.Phone,
		UserType:  profile.UserType,
		Address:   profile.Address,
		Cart:      profile.Cart,
		Orders:    profile.Orders,
	}
	return &userProfile, nil
}

func (s *UserService) UpdateProfile(id int, input *dto.ProfileInput) error {
	//Getting current user
	user, err := s.Repo.FindUserbyID(id)
	if err != nil {
		return err
	}

	//firstname,lastname update check
	if input.FirstName != "" {
		user.FirstName = input.FirstName
	}
	if input.LastName != "" {
		user.LastName = input.LastName
	}

	//update user with current firstname,lastname
	_, err = s.Repo.UpdateUser(id, user)
	if err != nil {
		return err
	}

	//update profile address
	if err := s.Repo.UpdateProfile(&domain.Address{
		AddressLine1: input.AddressInput.AddressLine1,
		AddressLine2: input.AddressInput.AddressLine2,
		City:         user.Address.City,
		Country:      input.AddressInput.Country,
		PostCode:     input.AddressInput.PostCode,
	}); err != nil {
		return errors.New("profile addres updation failed")
	}

	return nil
}

func (s *UserService) BecomeSeller(id int, input dto.SellerInput) (string, error) {
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

func (s *UserService) FindCart(id uint) ([]*domain.Cart, error) {

	cart, err := s.Repo.FindCartItems(int(id))
	if err != nil {
		log.Printf("cart fetching failed at service layer %v", err)
		return nil, errors.New("something went wrong while fetching cart")
	}

	return cart, nil
}

func (s *UserService) CreateCart(input *dto.CreateCartRequest, u domain.User) ([]*domain.Cart, error) {
	//Check if cart exist
	cart, _ := s.Repo.FindCartItem(u.ID, int(input.ProductId))
	if cart.ID > 0 {
		if input.ProductId == 0 {
			return nil, errors.New("invalid product id")
		}

		if input.Qty < 1 {
			//delete the cart
			if err := s.Repo.DeleteCartItemByid(u.ID); err != nil {
				log.Printf("error on deleting cart item %v", err)
				return nil, errors.New("error on deleting cart item")
			}
		} else {
			//update the cart
			cart.Qty = int(input.Qty)
			if err := s.Repo.UpdateCart(*cart); err != nil {
				//log error
				return nil, errors.New("updating cart failed")
			}
			return s.Repo.FindCartItems(u.ID)
		}
	} else {
		//check if product exist for creating cart
		product, err := s.CRepo.FindProductById(int(input.ProductId))
		if err != nil {
			return nil, errors.New("product not found to create cart item")
		}

		//create cart
		err = s.Repo.CreateCart(domain.Cart{
			UserId:    u.ID,
			ProductId: int(input.ProductId),
			Name:      product.Name,
			ImageUrl:  product.ImageUrl,
			Price:     product.Price,
			Qty:       int(input.Qty),
			SellerId:  product.UserId,
		})

		if err != nil {
			log.Printf("cart creation error-%v", err)
			return nil, errors.New("error on creating cart item")
		}

	}

	return s.Repo.FindCartItems(u.ID)
}

func (s *UserService) CreateOrder(u domain.User) (int, error) {

	//Get cart items of current user
	cartItems, err := s.Repo.FindCartItems(u.ID)
	if err != nil {
		return 0, errors.New("could not find cart items")
	}

	//if above function return object with no item
	if len(cartItems) == 0 {
		return 0, errors.New("no items found")
	}

	//find success payment status
	paymentId := "PAY12345"
	txnId := "TXN12345"
	orderRef, _ := helper.RandomNumbers(8)

	//create order with generated OrderNumber
	var amount float64
	var orderItems []domain.OrderItem

	for _, item := range cartItems {
		amount += item.Price * float64(item.Qty)
		orderItems = append(orderItems, domain.OrderItem{
			ProductId: item.ProductId,
			Name:      item.Name,
			Price:     item.Price,
			Qty:       item.Qty,
			ImageUrl:  item.ImageUrl,
			SellerId:  item.SellerId,
		})
	}
	order := domain.Order{
		UserId:         uint(u.ID),
		PaymentId:      paymentId,
		TransactionId:  txnId,
		OrderRefNumber: orderRef,
		Amount:         amount,
		Items:          orderItems,
	}
	if err := s.Repo.CreateOrder(order); err != nil {
		return 0, err
	}

	//Delete items from cart after order success
	if err := s.Repo.DeleteCartItems(u.ID); err != nil {
		return 0, err
	}

	return orderRef, nil
}

func (s *UserService) GetOrders(userId int) ([]*domain.Order, error) {

	orders, err := s.Repo.FindOrders(userId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *UserService) GetOrderById(orderId uint, userId int) (*domain.Order, error) {

	order, err := s.Repo.FindOrderById(int(orderId), userId)
	if err != nil {
		return nil, err
	}
	return order, nil
}
