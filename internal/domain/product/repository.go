package product

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	Update(ctx context.Context, id uuid.UUID, product *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}
