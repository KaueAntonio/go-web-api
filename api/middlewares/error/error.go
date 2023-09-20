package erroService

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
)

func Error(w http.ResponseWriter, r *http.Request) {

	msg := chi.URLParam(r, "msg")

	codeParam := chi.URLParam(r, "code")

	code, err := strconv.ParseInt(codeParam, 0, 64)

	if err != nil {
		panic("Erro ao decodificar código de erro!")
	}

	w.WriteHeader(int(code))

	http.Error(w, decodeMsg(msg, w), int(code))
}

func decodeMsg(msg string, w http.ResponseWriter) (m string) {
	unescapedMessage, err := url.QueryUnescape(msg)

	if err != nil {
		http.Error(w, "Erro ao decodificar a mensagem", http.StatusInternalServerError)
		return
	}

	return unescapedMessage
}
