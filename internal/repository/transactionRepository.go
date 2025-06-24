package repository

import (
	"go-ecommerce-app/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreatePayment(payment *domain.Payment) error
	FindOrders(userId int) ([]*domain.OrderItem, error)
	FindOrderById(orderI int, userId int) (*domain.Order, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (t *transactionRepo) CreatePayment(payment *domain.Payment) error {
	panic("unimplemented")
}

func (t *transactionRepo) FindOrderById(orderI int, userId int) (*domain.Order, error) {
	panic("unimplemented")
}

func (t *transactionRepo) FindOrders(userId int) ([]*domain.OrderItem, error) {
	panic("unimplemented")
}
