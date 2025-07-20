package repositories

import (
	"context"
	"curso-go/internal/categories"

	"github.com/jmoiron/sqlx"
)

type GetRepository struct {
	db *sqlx.DB
}

func NewGetRepository(db *sqlx.DB) GetRepository {
	return GetRepository{
		db: db,
	}
}

func (repo *GetRepository) GetCategory(ctx context.Context, categoryID string) (categories.Category, error) {
	var model CategoryModel

	err := repo.db.SelectContext(ctx, &model, "SELECT * FROM categories WHERE id = ?", categoryID)

	if err != nil {
		return categories.Category{}, err
	}

	return model.ToDomain(), nil
}
