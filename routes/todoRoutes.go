package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/maruti-jogdand/todo-application/controllers"
)

func GetNewRouter() *mux.Router {
	return mux.NewRouter()
}

func RegisterRoutes(muxRouter *mux.Router, route controllers.ItemAPI) {
	muxRouter.Handle("/", http.HandlerFunc(controllers.Welcome))

	// Item Route
	muxRouter.Handle("/items/create", http.HandlerFunc(route.CreateItem)).Methods("POST")
	muxRouter.Handle("/items/update/{id}", http.HandlerFunc(route.UpdateItem)).Methods("PUT")
	muxRouter.Handle("/items/delete/{id}", http.HandlerFunc(route.DeleteItem)).Methods("DELETE")
	muxRouter.Handle("/items/{id}", http.HandlerFunc(route.GetItem)).Methods("GET")
	muxRouter.Handle("/items/", http.HandlerFunc(route.GetItems)).Methods("GET")
}
