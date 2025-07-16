package categories

type Service struct {
	createRepository CreateCategoryGateway
}

func New(opts ...ServiceOption) Service {
	s := Service{}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

type ServiceOption func(*Service)

func WithCreateRepository(createRepository CreateCategoryGateway) ServiceOption {
	return func(s *Service) {
		s.createRepository = createRepository
	}
}
