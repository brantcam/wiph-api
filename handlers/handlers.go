package handlers

import (
	"encoding/json"
	"net/http"
)

// HelloWorld ...
func HelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSONResponse(w, http.StatusOK, "I love you, mom!")
	}
}

func writeJSONResponse(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Cache-Control", "no-cache")
	buf, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(buf)
}
