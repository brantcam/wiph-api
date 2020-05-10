package main

import (
//	"context"
	"log"
	"net/http"
//	"wiph-api/config"
	"wiph-api/router"
//	"wiph-api/store/postgres"
//	"wiph-api/users"
)

func main() {
//	ctx := context.Background()
//	c, err := config.LoadConfigFromEnv()
//	if err != nil {
//		log.Fatal(err)
//	}
//	db, err := postgres.New(ctx, c.Postgres)
//	if err != nil {
//		log.Fatal(err)
//	}
//	userOps := &users.UserOps{Pg: db}
//	handler := router.New(router.Options{
//		Users: userOps,
//	})
	handler := router.New()

	srv := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	log.Printf("listening on port %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
