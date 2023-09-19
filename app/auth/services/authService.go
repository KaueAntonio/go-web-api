package authService

import (
	"net/http"
	errors "web/app/api/middlewares/error"
	"web/app/api/middlewares/result"
	"web/app/app/auth/models"
	authRepository "web/app/app/auth/repositories"
	bodyEncoding "web/app/app/json"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var p models.Pessoa

	p = bodyEncoding.DecodeBodyPerson(p, w, r)

	id, err := authRepository.Login(p)

	if err != nil || id <= 0 {
		errors.HandleError(w, r, "Usuário não encontrado", 3)
	}

	result.GetResult(id, w)
}
