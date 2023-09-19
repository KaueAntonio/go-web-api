package controllers

import (
	authService "web/app/app/auth/services"

	"github.com/go-chi/chi"
)

func AddAuthController(r *chi.Mux) {
	r.Post("/login", authService.Login)
}
