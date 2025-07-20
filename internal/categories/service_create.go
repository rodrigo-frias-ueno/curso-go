package categories

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

// CreateCategory creates a new category with the provided candidate data.
// It assigns a new UUID, sets the status to active, and timestamps for creation and update.
// The candidate is validated before being persisted using the repository.
// Returns the created Category or an error if validation or persistence fails.
func (s *Service) CreateCategory(ctx context.Context, candidate Category) (Category, error) {
	candidate.ID = uuid.NewString()
	candidate.Status = StatusActive
	candidate.CreatedAt = time.Now()
	candidate.UpdatedAt = candidate.CreatedAt

	if err := candidate.validate(); err != nil {
		return Category{}, err
	}

	if err := s.createRepository.Create(ctx, &candidate); err != nil {
		slog.ErrorContext(ctx, "failed to create category", "error", err, "category", candidate)

		return Category{}, ErrInternal
	}

	return candidate, nil
}
