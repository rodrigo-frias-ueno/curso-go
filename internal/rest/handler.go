package rest

type Handler struct {
	service CategoryServiceGateway
}

func NewHandler(service CategoryServiceGateway) *Handler {
	return &Handler{
		service: service,
	}
}
