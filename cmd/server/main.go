package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/AI1411/connectpj/gen/proto/user/v1/userv1connect"

	"github.com/AI1411/connectpj/internal/env"
	"github.com/AI1411/connectpj/internal/infra/db"
	"github.com/AI1411/connectpj/internal/server"
)

func main() {
	mux := http.NewServeMux()
	e, err := env.NewValue()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
		return
	}
	dbClient := db.NewClient(&e.DB)
	userServer := server.NewUserServer(dbClient)
	path, handler := userv1connect.NewUserServiceHandler(userServer)
	log.Printf("path: %+v", path)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
