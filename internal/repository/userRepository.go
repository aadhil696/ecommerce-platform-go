package repository

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserbyID(id int) (domain.User, error)
	UpdateUser(id int, usr domain.User) (domain.User, error)
	AddBankAccount(e domain.BankAccount) error
}

type userRepository struct {
	db *gorm.DB
}

// Addind bank account for seller feature
func (r *userRepository) AddBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) CreateUser(usr domain.User) (domain.User, error) {

	result := r.db.Create(&usr) // Just use the input directly
	if result.Error != nil {
		log.Println("User creation failed due to", result.Error)
		return domain.User{}, result.Error
	}

	return usr, nil
}

func (r *userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	//Query the database for a user with the given email
	result := r.db.Where("email=?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("User with email %s not found", email)
			return domain.User{}, fmt.Errorf("user not found")
		}

		log.Printf("Error finding user :%v", result.Error)
		return domain.User{}, fmt.Errorf("database error: %w", result.Error)
	}

	return user, nil
}

func (r *userRepository) FindUserbyID(id int) (domain.User, error) {
	var user domain.User

	result := r.db.Where("id=?", id).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("User with id %d not found", id)
			return domain.User{}, fmt.Errorf("user not found")
		}

		log.Printf("Error finding user :%v", result.Error)
		return domain.User{}, fmt.Errorf("database error: %w", result.Error)
	}

	return user, nil
}

func (r *userRepository) UpdateUser(id int, usr domain.User) (domain.User, error) {
	var existingUser domain.User

	//check if user with id exist
	if err := r.db.First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, fmt.Errorf("user with ID %d not found", id)
		}

		return domain.User{}, fmt.Errorf("database error: %w", err)
	}

	//Update the user fields
	result := r.db.Model(&domain.User{}).Where("id=?", id).Updates(usr)
	if result.Error != nil {
		log.Printf("Error updating user: %v", result.Error)
		return domain.User{}, fmt.Errorf("failed to update user: %w", result.Error)
	}

	//Verify the update was successfull
	if result.RowsAffected == 0 {
		return domain.User{}, fmt.Errorf("no user was updated")
	}

	//Retrieve and return the updated user
	var updatedUser domain.User
	if err := r.db.First(&updatedUser, id).Error; err != nil {
		return domain.User{}, fmt.Errorf("failed to fetch updated user: %w", err)
	}

	return updatedUser, nil

}
