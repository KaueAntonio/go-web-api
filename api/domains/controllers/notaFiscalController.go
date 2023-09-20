package controllers

import (
	notaFiscalService "web/app/app/notaFiscal/services"

	"github.com/go-chi/chi"
)

func AddNotaFiscalController(r *chi.Mux) {
	r.Get("/notas-fiscais", notaFiscalService.GetNotasFiscais)
}
