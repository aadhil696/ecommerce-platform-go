package repository

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"log"

	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateCategory(e *domain.Category) (*domain.Category, error)
	FindCategories() ([]domain.Category, error)
	FindCategoryById(id int) (*domain.Category, error)
	EditCategory(id int, e *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) (err error)

	CreateProduct(prdct *domain.Product) (*domain.Product, error)
	FindProduct() ([]*domain.Product, error)
	FindProductById(id int) (*domain.Product, error)
	FindSellerProducts(id int) ([]*domain.Product, error)
	UpdateProduct(prdct *domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
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
func (c *catalogRepository) CreateCategory(e *domain.Category) (*domain.Category, error) {
	result := c.db.Create(&e)
	if result.Error != nil {
		log.Printf("category creation failed due to db error- %v", result.Error)
		return nil, fmt.Errorf("category creation failed at db level %v", result.Error)
	}

	return e, nil
}

// DeleteCategory implements CatalogRepository.
func (c *catalogRepository) DeleteCategory(id int) (err error) {

	result := c.db.Delete(&domain.Category{}, id)
	if result.Error != nil {
		log.Printf("category deletion failed due to db error- %v", err)
		return errors.New("cateogry deletion failed at db level")
	}

	return nil
}

// EditCategory implements CatalogRepository.
func (c *catalogRepository) EditCategory(id int, e *domain.Category) (*domain.Category, error) {

	err := c.db.Save(&e).Error
	if err != nil {
		log.Printf("Category editing failed at db level due to %v", err.Error())
		return &domain.Category{}, errors.New("category modification failed")
	}

	return e, nil
}

// FindCategories implements CatalogRepository.
func (c *catalogRepository) FindCategories() ([]domain.Category, error) {
	var categories []domain.Category

	result := c.db.Find(&categories)
	if result.Error != nil {
		log.Printf("Error fetching categories: %v", result.Error)
		return nil, result.Error
	}

	return categories, nil
}

// FindCategoryById implements CatalogRepository.
func (c *catalogRepository) FindCategoryById(id int) (*domain.Category, error) {
	var category domain.Category

	result := c.db.Where("id=?", id).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("category with id %d not found", id)
			return &domain.Category{}, fmt.Errorf("category not found")
		}

		log.Printf("Error finding category :%v", result.Error)
		return &domain.Category{}, fmt.Errorf("database error: %w", result.Error)
	}
	return &category, nil

}

func (c *catalogRepository) CreateProduct(prdct *domain.Product) (*domain.Product, error) {

	result := c.db.Model(&domain.Product{}).Create(&prdct)
	if result.Error != nil {
		log.Println("Product creation failed at db level", result.Error)
		return nil, result.Error
	}

	return prdct, nil
}

func (c *catalogRepository) FindProduct() ([]*domain.Product, error) {
	var allProducts []*domain.Product
	result := c.db.Find(&allProducts)
	if result.Error != nil {
		log.Println("product fetching failed at db level", result.Error)
		return nil, errors.New("fetching products failed due to some internal error")
	}

	return allProducts, nil
}

func (c *catalogRepository) FindProductById(id int) (*domain.Product, error) {
	var product *domain.Product
	result := c.db.Where("id=?", id).First(&product)
	if result.Error != nil {
		log.Println("Product fetching db error", result.Error)
		return nil, errors.New("product fetching failed-db error")
	}

	return product, nil
}

func (c *catalogRepository) FindSellerProducts(id int) ([]*domain.Product, error) {
	var prdcts []*domain.Product

	result := c.db.Where("userid=?", id).Find(&prdcts)
	if result.Error != nil {
		// log.Println("finding seller products db error", result.Error)
		return nil, fmt.Errorf("fetching sellers products failed due to -%s", result.Error)
	}

	return prdcts, nil
}

func (c *catalogRepository) UpdateProduct(prdct *domain.Product) (*domain.Product, error) {

	err := c.db.Save(&prdct).Error
	if err != nil {
		// log.Printf("product editing failed at db level due to %v", err.Error())
		return &domain.Product{}, fmt.Errorf("product updation failed due to-%s", err.Error())
	}

	return prdct, nil
}

func (c *catalogRepository) DeleteProduct(id int) error {

	result := c.db.Delete(&domain.Product{}, id)
	if result.Error != nil {
		// log.Println("Product deletion failed at db level", result.Error)
		return errors.New("product deletion failed due to some internal error")
	}

	return nil
}
