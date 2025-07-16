package categories

type CategoryStatus string

func (s CategoryStatus) validate() error {
	if s != StatusActive && s != StatusInactive {
		return ErrCategoryStatusInvalid
	}

	return nil
}

const (
	StatusActive   CategoryStatus = "active"
	StatusInactive CategoryStatus = "inactive"
)
