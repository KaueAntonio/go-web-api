package result

import (
	"encoding/json"
	"net/http"
)

func GetResult(r string, w http.ResponseWriter) {

	resp := map[string]any{
		"Data": r,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
