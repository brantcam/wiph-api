package router

import (
	"net/http"
	"wiph-api/handlers"

	"github.com/gorilla/mux"
)

// New ...
func New() *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/hello-world").Handler(handlers.HelloWorld())
	return r
}
