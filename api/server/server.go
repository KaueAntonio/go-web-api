package server

import (
	"fmt"
	"net/http"
	"web/app/api/config"
	"web/app/api/routers"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func ExposeServer() {
	err := config.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	routers.ExposeControllers(r)

	fmt.Println("Server Iniciado Com Sucesso\nUrl: http://localhost:" + config.GetServerPort())

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

	http.Handle("swagger", httpSwagger.Handler(
		httpSwagger.URL("localhost:"+config.GetServerPort()+"/swagger/doc.json"),
	))
}
