package categories

import "context"

type CreateCategoryGateway interface {
	Create(ctx context.Context, category *Category) error
}
