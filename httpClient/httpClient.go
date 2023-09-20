package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	config "web/app/api/config"
)

func Get(url string) string {
	fullUrl := config.GetBaseUrl() + url

	client := &http.Client{}

	req, err := http.NewRequest("GET", fullUrl, nil)

	//req.Header.Add("chave-api-dados ", config.GetChaveApi())

	if err != nil {
		fmt.Println("Erro ao criar a solicitação:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
	}

	decodedBody := string(body)

	return decodedBody

}
