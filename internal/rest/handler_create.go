package rest

import (
	"curso-go/internal/categories"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ueno-tecnologia-org/go-core/pkg/web"
)

// CreateCategory creates a new category.
//
// @Summary      Create a new category
// @Description  Creates a new category with the provided information.
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      CreateCategoryRequest  true  "Category to create"
// @Success      201       {object}  CategoryApiResponse
// @Failure      400       {object}  error  "Example: {\"code\": \"Bad Request\", \"message\": \"Invalid request body\"}"
// @Failure      500       {object}  error  "Example: {\"code\": \"Internal Server Error\", \"message\": \"Invalid request body\"}"
// @Router       /api/v1/categories [post]
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
