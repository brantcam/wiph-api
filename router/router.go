package router

import (
	"net/http"
	"wiph-api/handlers"
	"wiph-api/users"

	"github.com/gorilla/mux"
)

// Options ...
type Options struct {
	Users users.Repo
}

// New will take (opts Options) but it was taken out for simplicities sake
func New() *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/hello-world").Handler(handlers.HelloWorld())
//	r.Methods(http.MethodGet).Path("/get-users").Handler(handlers.GetUsers(opts.Users))
	return r
}
