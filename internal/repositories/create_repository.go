package repositories

import (
	"context"
	"curso-go/internal/categories"

	"github.com/jmoiron/sqlx"
)

type CreateRepository struct {
	db *sqlx.DB
}

func NewCreateRepository(db *sqlx.DB) CreateRepository {
	return CreateRepository{
		db: db,
	}
}

func (repo *CreateRepository) Create(ctx context.Context, category *categories.Category) error {
	model := FromDomain(*category)

	query := `INSERT INTO categories (id, name, description, parent_id, status, created_at, updated_at)
	VALUES (:id, :name, :description, :parent_id, :status, :created_at, :updated_at);`

	_, err := repo.db.ExecContext(ctx, query, model)

	if err != nil {
		return err
	}

	return nil
}
