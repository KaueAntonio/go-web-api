package erroService

import (
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
)

func Error404(w http.ResponseWriter, r *http.Request) {

	msg := chi.URLParam(r, "msg")

	w.WriteHeader(http.StatusNotFound)

	http.Error(w, decodeMsg(msg, w), 404)
}

func Error500(w http.ResponseWriter, r *http.Request) {

	msg := chi.URLParam(r, "msg")

	w.WriteHeader(http.StatusInternalServerError)

	http.Error(w, decodeMsg(msg, w), 500)
}

func Error401(w http.ResponseWriter, r *http.Request) {

	msg := chi.URLParam(r, "msg")

	w.WriteHeader(http.StatusUnauthorized)

	http.Error(w, decodeMsg(msg, w), 401)
}

func decodeMsg(msg string, w http.ResponseWriter) (m string) {
	unescapedMessage, err := url.QueryUnescape(msg)

	if err != nil {
		http.Error(w, "Erro ao decodificar a mensagem", http.StatusInternalServerError)
		return
	}

	return unescapedMessage
}
