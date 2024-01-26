package main

import (
	"log"

	"github.com/luist1228/go-htmx-examples/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)

	}

	err = server.Start(":8080")

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
