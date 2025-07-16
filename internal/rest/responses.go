package rest

import (
	"time"
)

type CategoryApiResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    string    `json:"parent_id"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
