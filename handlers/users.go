package handlers

import (
	"net/http"
	"wiph-api/users"
)

// GetUsers ...
func GetUsers(u users.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := u.GetUsers()
		if err != nil {
			writeJSONResponse(w, http.StatusServiceUnavailable, err)
			return
		}
		writeJSONResponse(w, http.StatusOK, u)
	}
}
