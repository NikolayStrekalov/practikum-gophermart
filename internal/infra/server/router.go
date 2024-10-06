package server

import (
	"github.com/NikolayStrekalov/practicum-gophermart/internal/infra/server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const compressionLevel = 5

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Compress(compressionLevel, "text/html", "application/json"))

	r.Method("POST", "/api/user/register",
		middleware.AllowContentType("application/json")(handlers.Register))
	return r
}
