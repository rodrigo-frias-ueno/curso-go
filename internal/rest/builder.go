package rest

import "curso-go/internal/categories"

type ResponseBuilder categories.Category

func (r ResponseBuilder) build() CategoryApiResponse {
	return CategoryApiResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		ParentID:    r.ParentID,
		Status:      string(r.Status),
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
