package erroService

import (
	"fmt"
	"net/http"
)

func HandleError(w http.ResponseWriter, r *http.Request, msg string, code int) {
	switch code {
	case 1:
		code = 404
	case 2:
		code = 500
	case 3:
		code = 401
	default:
		code = 500
	}

	http.Redirect(w, r, fmt.Sprintf("/error/%d/%s", code, msg), http.StatusFound)
}
