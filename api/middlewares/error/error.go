package erroService

import (
	"fmt"
	"net/http"
)

func Error404(w http.ResponseWriter, r *http.Request) {

	msg := r.URL.Query().Get("msg")

	if msg != "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro %s - Página não encontrada", msg)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Página não encontrada")
	}

	http.Error(w, msg, 404)
}

func Error500(w http.ResponseWriter, r *http.Request) {

	msg := r.URL.Query().Get("msg")

	if msg != "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro %s - Página não encontrada", msg)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Página não encontrada")
	}

	http.Error(w, msg, 500)
}

func Error401(w http.ResponseWriter, r *http.Request) {

	msg := r.URL.Query().Get("msg")

	if msg != "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro %s - Página não encontrada", msg)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Página não encontrada")
	}

	http.Error(w, msg, 401)
}
