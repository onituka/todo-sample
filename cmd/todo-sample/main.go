package main

import (
	"log"

	"github.com/onituka/todo-sample/infrastructure/router"
)

func main() {
	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}
