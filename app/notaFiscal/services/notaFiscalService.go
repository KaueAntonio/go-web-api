package notaFiscalService

import (
	"net/http"
	"web/app/api/middlewares/result"
	httpclient "web/app/httpClient"
)

func GetNotasFiscais(w http.ResponseWriter, r *http.Request) {
	resp := httpclient.Get("notas-fiscais")

	result.GetResult(resp, w)
}
