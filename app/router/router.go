package router

import (
	"github.com/go-chi/chi"
	"go-with-docker/app/requestLog"
	"go-with-docker/app/server"
)

func New(a *server.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", requestLog.NewHandler(a.HandleIndex, l))
	return r
}
