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

func (s *CatalogService) UpdateCategories(id int, input *dto.CreateCategoryRequest) (domain.Category, error) {

	currentCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		log.Println("No category found with specific id ", err)
		return domain.Category{}, errors.New("no category found with specific id")
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
		return domain.Category{}, errors.New("updating category failed due to some internal error")
	}
	return updatedCat, nil
}

func (s *CatalogService) GetCategoryById(id int) (domain.Category, error) {

	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		log.Printf("Category fetching with id failed at service layer due to %v", err)
		return domain.Category{}, errors.New("category fetching failed due to internal error")
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

func (s *CatalogService) CreateProduct(input *dto.CreateProductRequest) (*domain.Product, error) {
	return nil, nil
}

func (s *CatalogService) FindProduct(input *dto.CreateProductRequest) (*domain.Product, error) {
	return nil, nil
}
