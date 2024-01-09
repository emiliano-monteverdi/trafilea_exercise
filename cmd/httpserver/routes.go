package httpserver

import (
	"github.com/Trafilea/cmd/dependencies"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

func Routes(router *chi.Mux, d dependencies.Definition) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(100 * time.Millisecond))

	router.Route("/numbers", func(r chi.Router) {
		r.Post("/", d.NumberHandler.Create)                // POST /numbers
		r.Get("/{id}", d.NumberHandler.Get)                // GET  /numbers/123
		r.Get("/bulk/value", d.NumberHandler.BulkGetValue) // GET  /numbers/bulk/value
		r.Get("/bulk/type", d.NumberHandler.BulkGetType)   // GET  /number/bulk/type
	})
}
