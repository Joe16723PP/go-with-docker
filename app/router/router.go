package router

import (
	"github.com/go-chi/chi"
	"go-with-docker/app/server"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.MethodFunc("GET", "/", server.HandleIndex)
	return r
}
