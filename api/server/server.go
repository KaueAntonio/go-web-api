package server

import (
	"fmt"
	"net/http"
	"web/app/api/config"
	"web/app/api/routers"

	"github.com/go-chi/chi"
)

func Expose() {
	err := config.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	routers.ExposeControllers(r)

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}
