package categories

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateCategory(ctx context.Context, candidate Category) (Category, error) {
	candidate.ID = uuid.NewString()
	candidate.Status = StatusActive
	candidate.CreatedAt = time.Now()
	candidate.UpdatedAt = candidate.CreatedAt

	if err := candidate.validate(); err != nil {
		return Category{}, err
	}

	if err := s.createRepository.Create(ctx, &candidate); err != nil {
		return Category{}, err
	}

	return candidate, nil
}
