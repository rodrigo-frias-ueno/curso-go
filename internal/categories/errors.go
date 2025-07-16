package categories

type CategoryError string

func (e CategoryError) Error() string {
	return string(e)
}

const (
	// ErrCategoryIDInvalid is returned when the category ID is empty.
	ErrCategoryIDInvalid CategoryError = "invalid category: ID cannot be empty"
	// ErrCategoryNameInvalid is returned when the category name is empty.
	ErrCategoryNameInvalid CategoryError = "invalid category: category name cannot be empty"
	// ErrCategoryStatusInvalid is returned when the category status is not valid.
	ErrCategoryStatusInvalid CategoryError = "invalid category: status must be a valid value"

	// ErrInternal is returned when an internal error occurs
	ErrInternal CategoryError = "internal error while processing category"
)
