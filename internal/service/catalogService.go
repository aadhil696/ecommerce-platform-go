package service

import (
	"errors"
	"fmt"
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"log"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config configs.AppConfig
}

// Category Implementation
func (s *CatalogService) CreateCategories(input *dto.CreateCategoryRequest) (*domain.Category, error) {

	cat, err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		ParentId:     input.ParentId,
		DisplayOrder: input.DisplayOrder,
	})

	if err != nil {
		log.Printf("category creation failed at service layer due to %v", err)
		return nil, fmt.Errorf("category creation failed due to %v", err)
	}

	return cat, nil
}

func (s *CatalogService) UpdateCategories(id int, input *dto.CreateCategoryRequest) (*domain.Category, error) {

	currentCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		log.Println("No category found with specific id ", err)
		return &domain.Category{}, errors.New("no category found with specific id")
	}

	if len(input.Name) > 0 {
		currentCat.Name = input.Name
	}

	if input.ParentId > 0 {
		currentCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		currentCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		currentCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(id, currentCat)
	if err != nil {
		log.Println("Updating category failed at service layer", err)
		return &domain.Category{}, err
	}
	return updatedCat, nil
}

func (s *CatalogService) GetCategoryById(id int) (*domain.Category, error) {

	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		log.Printf("Category fetching with id failed at service layer due to %v", err)
		return &domain.Category{}, errors.New("category fetching failed due to internal error")
	}
	return cat, err
}

func (s *CatalogService) GetCategories() ([]*domain.Category, error) {

	result, err := s.Repo.FindCategories()
	if err != nil {
		log.Printf("All category fetching failed at service layer due to %v", err)
		return nil, errors.New("category list fetching failed due to some internal error")
	}

	var allcategories []*domain.Category
	for _, category := range result {
		allcategories = append(allcategories, &domain.Category{
			ID:           category.ID,
			Name:         category.Name,
			ImageUrl:     category.ImageUrl,
			ParentId:     category.ParentId,
			Products:     category.Products,
			DisplayOrder: category.DisplayOrder,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
		})
	}

	return allcategories, nil
}

func (s *CatalogService) DeleteCategories(id int) error {

	_, err := s.Repo.FindCategoryById(id)
	if err != nil {
		log.Println("no category found with specific id ", err)
		return fmt.Errorf("no category found with id %d", id)
	}

	if err := s.Repo.DeleteCategory(id); err != nil {
		log.Printf("Category deletion failed at service layer due to-%v", err)
		return errors.New("category deletion failed due to internal error")
	}
	return nil
}

// Product Implementation
func (s *CatalogService) CreateProduct(id int, input *dto.CreateProductRequest) (*domain.Product, error) {

	prdct, err := s.Repo.CreateProduct(&domain.Product{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		UserId:      id,
		ImageUrl:    input.ImageUrl,
		CategoryID:  input.CategoryID,
		Stock:       input.Stock,
	})
	if err != nil {
		log.Println("product creation service layer error", err)
		return nil, err
	}
	return prdct, nil
}

func (s *CatalogService) GetAllProducts() ([]*domain.Product, error) {

	allProducts, err := s.Repo.FindProduct()
	if err != nil {
		log.Println("product fetching failed at service layer", err)
		return nil, err
	}

	// var allProducts []*domain.Product
	// for _, product := range result {
	// 	allProducts = append(allProducts, &domain.Product{
	// 		ID:          product.ID,
	// 		Name:        product.Name,
	// 		Price:       product.Price,
	// 		Description: product.Description,
	// 		ImageUrl:    product.ImageUrl,
	// 		CategoryID:  product.CategoryID,
	// 		Stock:       product.Stock,
	// 	})
	// }
	return allProducts, nil
}

func (s *CatalogService) FindProductById(id int) (*domain.Product, error) {

	prdct, err := s.Repo.FindProductById(id)
	if err != nil {
		log.Println("product not found,service layer", err)
		return nil, err
	}

	return prdct, err
}

func (s *CatalogService) FindSellerProducts(id int) ([]*domain.Product, error) {

	sellersProducts, err := s.Repo.FindSellerProducts(id)
	if err != nil {
		log.Println("seller products fetching failed, service layer", err)
		return nil, err
	}

	// var sellerProducts []*domain.Product
	// for _, product := range result {
	// 	sellerProducts = append(sellerProducts, &domain.Product{
	// 		ID:          product.ID,
	// 		Name:        product.Name,
	// 		Price:       product.Price,
	// 		ImageUrl:    product.ImageUrl,
	// 		Description: product.Description,
	// 		Stock:       product.Stock,
	// 		UserId:      id,
	// 		CategoryID:  product.CategoryID,
	// 	})
	// }

	return sellersProducts, nil
}

func (s *CatalogService) UpdateProduct(id int, input *dto.CreateProductRequest, user *domain.User) (*domain.Product, error) {

	currentPrdct, err := s.Repo.FindProductById(id)
	if err != nil {
		log.Println("no product with id ", err)
		return &domain.Product{}, err
	}

	if currentPrdct.UserId != user.ID {
		return &domain.Product{}, errors.New("sorry, the product does not belongs to your stock")
	}

	//Update the current product field with non empty input data
	if len(input.Name) > 0 {
		currentPrdct.Name = input.Name
	}

	if input.Price > 0 {
		currentPrdct.Price = input.Price
	}

	if len(input.Description) > 0 {
		currentPrdct.Description = input.Description
	}

	if len(input.ImageUrl) > 0 {
		currentPrdct.ImageUrl = input.ImageUrl
	}

	if input.CategoryID > 0 {
		currentPrdct.CategoryID = input.CategoryID
	}

	if input.Stock > 0 {
		currentPrdct.Stock = input.Stock
	}

	updatedPrdct, err := s.Repo.UpdateProduct(currentPrdct)
	if err != nil {
		log.Println("product updation failed, service layer", err)
		return nil, err
	}

	return updatedPrdct, nil
}

func (s *CatalogService) StockUpdate(id int, input *dto.UpdateStockRequest, user domain.User) (*domain.Product, error) {

	prdct, err := s.Repo.FindProductById(id)
	if err != nil {
		log.Println("no product found with specific id", err)
		return &domain.Product{}, err
	}

	if prdct.UserId != user.ID {
		return &domain.Product{}, errors.New("sorry, the product does not belongs to your stock")
	}

	if input.Stock == prdct.Stock {
		return &domain.Product{}, errors.New("same stock quantity exist in storage")
	} else {
		prdct.Stock = input.Stock
	}

	updatedprct, err := s.Repo.UpdateProduct(prdct)
	if err != nil {
		log.Println("stock updation failed,service layer", err)
		return &domain.Product{}, err
	}

	return updatedprct, nil
}

func (s *CatalogService) DeleteProduct(id int) error {

	_, err := s.Repo.FindProductById(id)
	if err != nil {
		return errors.New("product not found")
	}

	if err := s.Repo.DeleteProduct(id); err != nil {
		log.Println("product deletion failed", err)
		return err
	}

	return nil
}
