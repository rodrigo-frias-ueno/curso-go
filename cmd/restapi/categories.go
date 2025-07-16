package main

import (
	"curso-go/internal/categories"
	"curso-go/internal/repositories"
	"curso-go/internal/rest"

	"github.com/jmoiron/sqlx"
)

func buildCategoriesHandler(db *sqlx.DB) *rest.Handler {
	createCategoryRepository := repositories.NewCreateRepository(db)
	categoryService := categories.New(categories.WithCreateRepository(&createCategoryRepository))
	categoryHandler := rest.NewHandler(&categoryService)

	return categoryHandler
}
