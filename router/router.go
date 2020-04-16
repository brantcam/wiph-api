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

// New ...
func New(opts Options) *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/hello-world").Handler(handlers.HelloWorld())
	r.Methods(http.MethodGet).Path("/get-users").Handler(handlers.GetUsers(opts.Users))
	return r
}
