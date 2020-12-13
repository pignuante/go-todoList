package main

import (
	"github.com/pignuante/go-todoList/server"
	"log"
)

func main() {
	if err := server.ListenAndServe(server.Config{
		Address:     ":8080",
		DatabaseURL: "",
	}); err != nil {
		log.Fatalln(err)
	}
}
