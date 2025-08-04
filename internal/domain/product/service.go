package product

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.starlark.net/lib/time"
)

type Service struct {
	repo Repository
}

// NewService creates a new product service with the given repository.
// This service will handle business logic related to products.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateProduct creates a new product using the provided request data.
// CreateProduct is the Method on Service
func (s *Service) CreateProduct(ctx context.Context, req *CreateProductRequest) (*Product, error) {

	// Validate the request data
	if req.Name == "" || req.Price <= 0 || req.Stock < 0 {
		return nil, fmt.Errorf("invalid product data")
	}

	// Create a new product instance
	product := &Product{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Active:      true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := s.repo.Create(ctx, product)

	return product, err
}

// get single product
// GetProduct is the Method on Service
func (s *Service) GetProduct(ctx context.Context, id uuid.UUID) (*Product, error) {
	return s.repo.GetByID(ctx, id)
}

// get all products
func (s *Service) GetAllProducts(ctx context.Context) ([]*Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) UpdateProduct(ctx context.Context, id uuid.UUID, req *UpdateProductRequest) (*Product, error) {
	product, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		product.Name = *req.Name
	}

	if req.Description != nil {
		product.Name = *req.Description
	}

	if req.Price != nil {
		product.Price = *req.Price
	}

	if req.Stock != nil {
		product.Stock = *req.Stock
	}

	product.UpdatedAt = time.Now()

	err = s.repo.Update(ctx, id, product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	product, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return err
	}

	product.Active = false

	updateErr := s.repo.Update(ctx, id, product)

	return updateErr

}

func (s *Service) Perma
