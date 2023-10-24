package main

import (
	"log"
	"github.com/gogogo/todo-app"
	"github.com/gogogo/todo-app/pkg/handler"

)


func main() {
	handlers := new(handler.Handler)
    srv := new(todo.Server)
    if err := srv.Run( "8000", handlers.initRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
 }
}