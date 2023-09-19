package result

import (
	"encoding/json"
	"net/http"
)

func GetResult(r any, w http.ResponseWriter) {
	message := "Sucesso!"

	if r != nil {
		r = ""
	}

	resp := map[string]any{
		"Message": message,
		"Data":    r,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
