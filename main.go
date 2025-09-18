package main

import (
	"fmt"
	"httpserver/http"
	"httpserver/todo"
)

func main() {
	todoList := todo.NewList()
	httpHadlers := http.NewHTTPHandler(todoList)
	httpServer := http.NewHTTPServer(httpHadlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}
}
