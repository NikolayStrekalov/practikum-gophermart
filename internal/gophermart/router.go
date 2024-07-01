package gophermart

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const indexPath = "/"
const compressionLevel = 5
const messageInternalServerError = "InternalServerError"

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Compress(compressionLevel, "text/html", "application/json"))

	r.Get(indexPath, indexHandler)

	return r
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	if _, err := res.Write([]byte("Hello, Gophermart!")); err != nil {
		http.Error(res, messageInternalServerError, http.StatusInternalServerError)
	}
}
