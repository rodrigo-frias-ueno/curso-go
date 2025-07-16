package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/ueno-tecnologia-org/go-core/pkg/web"
)

func buildRouter(db *sqlx.DB) *web.Router {
	categoryHandler := buildCategoriesHandler(db)

	router := web.New()

	router.Post("/api/v1/categories", categoryHandler.CreateCategory)

	return router
}
