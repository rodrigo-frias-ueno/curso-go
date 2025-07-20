package main

import (
	"net/http"

	_ "curso-go/docs"

	"github.com/jmoiron/sqlx"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/ueno-tecnologia-org/go-core/pkg/web"
)

func buildRouter(db *sqlx.DB) *web.Router {
	router := web.New()
	router.Use(web.HeaderForwarder())
	initSwagger(router)

	categoryHandler := buildCategoriesHandler(db)

	router.Post("/api/v1/categories", categoryHandler.CreateCategory)

	return router
}

func initSwagger(router *web.Router) {
	router.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) error {
		httpSwagger.WrapHandler.ServeHTTP(w, r)
		return nil
	})
}
