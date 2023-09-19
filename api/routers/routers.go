package routers

import (
	c "web/app/api/domains/controllers"

	"github.com/go-chi/chi"
)

func ExposeControllers(r *chi.Mux) {

	c.AddAuthController(r)

	c.AddErrorController(r)
}
