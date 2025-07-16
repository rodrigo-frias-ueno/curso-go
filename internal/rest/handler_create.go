package rest

import (
	"curso-go/internal/categories"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ueno-tecnologia-org/go-core/pkg/web"
)

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) error {
	var request CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		slog.ErrorContext(r.Context(), "Failed to decode request body", "error", err)

		return web.NewError(http.StatusBadRequest, "Invalid request body")
	}

	createdCategory, err := h.service.CreateCategory(r.Context(), request.ToCategory())

	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to create category", "error", fmt.Errorf("%w", err))

		return h.buildApiError(err)
	}

	response := ResponseBuilder(createdCategory).build()

	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(response)
}

func (h *Handler) buildApiError(err error) error {
	if err != categories.ErrInternal {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	return web.NewError(http.StatusInternalServerError, "An internal error occurred")
}
