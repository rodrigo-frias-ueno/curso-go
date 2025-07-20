package categories

import "context"

//go:generate mockgen -source=gateways.go -destination=gateways_test.go -package=categories_test
type CreateCategoryGateway interface {
	Create(ctx context.Context, category *Category) error
}
