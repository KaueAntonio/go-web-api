package controllers

import (
	getServerError "web/app/api/middlewares/error"

	"github.com/go-chi/chi"
)

func AddErrorController(r *chi.Mux) {
	r.Get("/error/404/{msg}", getServerError.Error404)

	r.Get("/error/500/{msg}", getServerError.Error500)

	r.Get("/error/401/{msg}", getServerError.Error401)
}
