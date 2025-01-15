package service

import (
	"to-do-list/internal/domain/models"
	"to-do-list/internal/domain/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.repo.Create(product)
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) (*models.Product, error) {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
