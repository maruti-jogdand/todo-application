package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maruti-jogdand/todo-application/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemAPI struct {
	ItemHandler models.ItemsController
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	response := models.Message{
		Data: "welcome to todo app",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&response)
}

func (api ItemAPI) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	item.ItemId = primitive.NewObjectID().Hex()
	json.NewDecoder(r.Body).Decode(&item)
	if err := api.ItemHandler.Create(item); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Message{Data: err})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&models.Message{Data: "Item created"})
}

func (api ItemAPI) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	if err := api.ItemHandler.Update(id, item); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Message{Data: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&models.Message{Data: "Item Updated"})
}

func (api ItemAPI) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := api.ItemHandler.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&models.Message{Data: "Item Deleted"})
}

func (api ItemAPI) GetItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := api.ItemHandler.GetById(id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&item)
}

func (api ItemAPI) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := api.ItemHandler.GetAll()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(&models.Message{Data: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&items)
}
