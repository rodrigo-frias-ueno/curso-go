package rest

import "curso-go/internal/categories"

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id"`
}

func (req *CreateCategoryRequest) ToCategory() categories.Category {
	return categories.Category{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
	}
}
