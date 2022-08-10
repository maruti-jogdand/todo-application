package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/maruti-jogdand/todo-application/databases"
	"github.com/maruti-jogdand/todo-application/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemHandler struct {
	Collection *mongo.Collection
}

func NewItemHandler() *ItemHandler {
	return &ItemHandler{Collection: databases.NewItemCollection()}
}

func (db ItemHandler) Create(item models.Item) error {
	result, err := db.Collection.InsertOne(context.Background(), item)
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Println(result)
	return nil
}

func (db ItemHandler) Update(id string, item models.Item) error {
	fmt.Println(id)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", item.Status}}}}
	result, err := db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Println(result.UpsertedID)
	return nil
}

func (db ItemHandler) Delete(id string) error {
	result, err := db.Collection.DeleteOne(context.Background(), bson.D{{"_id", id}}, nil)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("Item Already deleted")
	}
	return nil
}

func (db ItemHandler) GetById(id string) (models.Item, error) {
	result := db.Collection.FindOne(context.Background(), bson.D{{"_id", id}})
	var item models.Item
	if err := result.Decode(&item); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (db ItemHandler) GetAll() ([]models.Item, error) {
	items := make([]models.Item, 0)
	cursor, err := db.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return items, err
	}

	cursor.All(context.Background(), &items)
	return items, nil
}
