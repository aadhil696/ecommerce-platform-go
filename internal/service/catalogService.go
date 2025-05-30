package service

import (
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

func (s *CatalogService) CreateCategories(input *dto.CreateCategoryRequest) error {

	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		ParentId:     input.ParentId,
		DisplayOrder: input.DisplayOrder,
	})

	if err != nil {
		log.Printf("category creation failed at service layer due to %v", err)
		return fmt.Errorf("category creation failed due to %v", err)
	}

	return nil
}

func (s *CatalogService) UpdateCategories(input *dto.CreateCategoryRequest) error {
	return nil
}

func (s *CatalogService) GetCategoryById(id uint) (any, error) {
	return nil, nil
}

func (s *CatalogService) GetCategories(input *dto.CreateCategoryRequest) (any, error) {
	return nil, nil
}

func (s *CatalogService) DeleteCategories(id uint) (any, error) {
	return nil, nil
}
