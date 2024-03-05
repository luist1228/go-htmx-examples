package main

import (
	"log"

	"github.com/luist1228/go-htmx-examples/api"
)

func main() {

	// todos := db.FillTodos()
	// util.Message("BEFORE")
	// todos.PrintTodos()

	// todos.Add("1")
	// util.Message("ADDED")
	// todos.PrintTodos()
	// todos.Add("2")
	// util.Message("ADDED")
	// todos.PrintTodos()
	// util.Message("ADDED")
	// todos.Add("3")
	// todos.PrintTodos()

	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)

	}

	err = server.Start(":5000")

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
