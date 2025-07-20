package rest

import "curso-go/internal/categories"

// CreateCategoryRequest represents the payload to create a new category.
//
// swagger:model CreateCategoryRequest
type CreateCategoryRequest struct {
	// Name of the category
	// required: true
	Name string `json:"name"`
	// Description of the category
	// required: false
	Description string `json:"description"`
	// Parent category ID
	// required: false
	ParentID string `json:"parent_id"`
}

func (req *CreateCategoryRequest) ToCategory() categories.Category {
	return categories.Category{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
	}
}
