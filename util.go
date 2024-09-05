package main

import (
	"encoding/json"
	"net/http"
)

func setJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
func WriteJSONResponse(w http.ResponseWriter, code int, data interface{}) {

	setJSONHeader(w)
	w.WriteHeader(code)
	switch x := data.(type) {
	case string:
		w.Write([]byte(x))
	case []byte:
		w.Write(x)
	default:
		err := json.NewEncoder(w).Encode(x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
