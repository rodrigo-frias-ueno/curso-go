package rest

import (
	"time"
)

// CategoryApiResponse represents the response for a category.
//
// swagger:model CategoryApiResponse
type CategoryApiResponse struct {
	// The unique identifier of the category
	// example: 123e4567-e89b-12d3-a456-426614174000
	ID string `json:"id"`
	// The name of the category
	// example: Electronics
	Name string `json:"name"`
	// The description of the category
	// example: Devices and gadgets
	Description string `json:"description"`
	// The parent category ID
	// example: 987e6543-e21b-12d3-a456-426614174111
	ParentID string `json:"parent_id"`
	// The status of the category
	// example: active
	Status string `json:"status,omitempty"`
	// The creation timestamp
	// example: 2024-06-01T12:00:00Z
	CreatedAt time.Time `json:"created_at"`
	// The last update timestamp
	// example: 2024-06-02T15:30:00Z
	UpdatedAt time.Time `json:"updated_at"`
}
