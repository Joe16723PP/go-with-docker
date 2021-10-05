package router

import (
	"github.com/go-chi/chi"
	"go-with-docker/app/requestLog"
	"go-with-docker/app/router/middleware"
	"go-with-docker/app/server"
)

func New(a *server.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	// Index route
	r.Method("GET", "/", requestLog.NewHandler(a.HandleIndex, l))

	// health check
	//r.Get("/health/liveness", server.HandleLive)
	r.Method("GET", "/health/readiness", requestLog.NewHandler(a.HandleReady, l))

	// grouping routes
	r.Route("api/v1", func(r chi.Router) {
		// attach middleware
		r.Use(middleware.ContentTypeJson)
		// Routes for books
		r.Method("GET", "/books", requestLog.NewHandler(a.HandleListBooks, l))
		r.Method("POST", "/books", requestLog.NewHandler(a.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", requestLog.NewHandler(a.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", requestLog.NewHandler(a.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", requestLog.NewHandler(a.HandleDeleteBook, l))
	})

	return r
}
