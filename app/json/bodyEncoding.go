package bodyEncoding

import (
	"encoding/json"
	"net/http"
	errors "web/app/api/middlewares/error"
	"web/app/app/auth/models"
)

func DecodeBodyPerson(b models.Pessoa, w http.ResponseWriter, r *http.Request) (_ models.Pessoa) {
	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		errors.HandleError(w, r, "Erro ao fazer decode do corpo da requisição", 2)
		return
	}

	return b
}
