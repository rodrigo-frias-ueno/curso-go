package rest

import (
	"context"
	"curso-go/internal/categories"
)

type CategoryServiceGateway interface {
	CreateCategory(ctx context.Context, candidate categories.Category) (categories.Category, error)
}
