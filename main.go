package main

import (
	"fmt"
	"net/http"

	"github.com/maruti-jogdand/todo-application/controllers"
	"github.com/maruti-jogdand/todo-application/handlers"
	"github.com/maruti-jogdand/todo-application/routes"
)

func main() {
	fmt.Println("Welcome to ToDo Application")

	router := routes.GetNewRouter()

	route := controllers.ItemAPI{
		ItemHandler: handlers.NewItemHandler(),
	}

	routes.RegisterRoutes(router, route)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
