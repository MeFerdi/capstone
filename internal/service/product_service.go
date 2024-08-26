package service

import (
	"capstone/internal/model"
	"capstone/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) CreateProduct(product model.Product) error {
	return s.repo.CreateProduct(product)
}
