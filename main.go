package main

import (
	"log"

	"github.com/lipandr/go-avito-trainee/internal"
)

func main() {
	cfg := internal.InitConfig()
	stg := internal.NewStorage()
	svc := internal.NewService(stg)

	server := internal.NewServer(cfg, svc)

	log.Fatal(server.ListenAndServe())

}
