package repository

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"log"

	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryById(id uint) (domain.Category, error)
	EditCategory(id uint, e *domain.Category) (*domain.Category, error)
	DeleteCategory(id uint) (err error)
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}

// CreateCategory implements CatalogRepository.
func (c *catalogRepository) CreateCategory(e *domain.Category) error {
	result := c.db.Create(&e)
	if result.Error != nil {
		log.Printf("User creation failed due to db error- %v", result.Error)
		return fmt.Errorf("category creation failed at db level %v", result.Error)
	}

	return nil
}

// DeleteCategory implements CatalogRepository.
func (c *catalogRepository) DeleteCategory(id uint) (err error) {

	if err := c.db.Delete(&domain.Category{}, id); err != nil {
		log.Printf("category deletion failed due to db error- %v", err)
		return errors.New("cateogry deletion failed at db level")
	}

	return nil
}

// EditCategory implements CatalogRepository.
func (c *catalogRepository) EditCategory(id uint, e *domain.Category) (*domain.Category, error) {

	err := c.db.Save(&e).Error
	if err != nil {
		log.Printf("Category editing failed at db level due to %v", err.Error())
		return &domain.Category{}, errors.New("category modification failed")
	}

	return e, nil
}

// FindCategories implements CatalogRepository.
func (c *catalogRepository) FindCategories() ([]*domain.Category, error) {
	var categories []*domain.Category

	if err := c.db.Find(&categories); err != nil {
		log.Printf("Something gone wrong while fetching all categories %v", err.Error)
		return nil, err.Error
	}

	return categories, nil
}

// FindCategoryById implements CatalogRepository.
func (c *catalogRepository) FindCategoryById(id uint) (domain.Category, error) {
	var category domain.Category

	if err := c.db.First(&category, id); err != nil {
		log.Printf("Category searching by id failed due to %v", err.Error)
		return domain.Category{}, err.Error
	}

	return category, nil

}
