package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepo
	Auth helper.Auth
}

func NewTransactionService(r repository.TransactionRepo, auth helper.Auth) *TransactionService {
	return &TransactionService{
		Repo: r,
		Auth: auth,
	}
}

func (s *TransactionService) GetOrders(u domain.User) ([]*domain.OrderItem, error) {

	orderItems, err := s.Repo.FindOrders(u.ID)
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (s *TransactionService) GetOrdersById(user domain.User, orderId int) (*domain.Order, error) {

	OrderDetails, err := s.Repo.FindOrderById(orderId, user.ID)
	if err != nil {
		return nil, err
	}

	return OrderDetails, nil
}
