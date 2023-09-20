package authService

import (
	"fmt"
	"net/http"
	errors "web/app/api/middlewares/error"
	"web/app/api/middlewares/result"
	"web/app/app/auth/models"
	authRepository "web/app/app/auth/repositories"
	bodyEncoding "web/app/app/json"
)

// @Summary Exemplo de rota Swagger
// @Description Esta rota é usada para fins de exemplo Swagger.
// @ID exemplo-rota-swagger
// @Produce json
// @Success 200
// @Router /exemplo [get]
func Login(w http.ResponseWriter, r *http.Request) {
	var p models.Pessoa

	p = bodyEncoding.DecodeBodyPerson(p, w, r)

	id, err := authRepository.Login(p)

	if err != nil || id <= 0 {
		errors.HandleError(w, r, "Usuário não encontrado", 3)
	}

	result.GetResult(fmt.Sprint(id), w)
}
