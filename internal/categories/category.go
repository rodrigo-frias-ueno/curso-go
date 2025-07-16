package categories

import "time"

type Category struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ParentID    string         `json:"parent_id"`
	Status      CategoryStatus `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (c Category) validate() error {
	if c.ID == "" {
		return ErrCategoryIDInvalid
	}

	if c.Name == "" {
		return ErrCategoryNameInvalid
	}

	if err := c.Status.validate(); err != nil {
		return err
	}

	return nil
}
