package controllers

import (
	getServerError "web/app/api/middlewares/error"

	"github.com/go-chi/chi"
)

func AddErrorController(r *chi.Mux) {
	r.Get("/error/{code}/{msg}", getServerError.Error)
}
