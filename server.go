package main

import (
	"net/http"
	"wiph-api/router"
)

func main() {
	handler := router.New()

	srv := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	srv.ListenAndServe()
}
