package repositories

import (
	"curso-go/internal/categories"
	"time"
)

type CategoryModel struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description *string   `db:"description"`
	ParentID    string    `db:"parent_id"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (model CategoryModel) ToDomain() categories.Category {
	return categories.Category{
		ID:          model.ID,
		Name:        model.Name,
		Description: *model.Description,
		ParentID:    model.ParentID,
		Status:      categories.CategoryStatus(model.Status),
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

}

func FromDomain(c categories.Category) CategoryModel {
	return CategoryModel{
		ID:          c.ID,
		Name:        c.Name,
		Description: &c.Description,
		ParentID:    c.ParentID,
		Status:      string(c.Status),
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
