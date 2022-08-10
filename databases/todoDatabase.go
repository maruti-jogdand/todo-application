package databases

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

const uri = "mongodb://localhost:27017"
const database = "todo-db"

func NewDbInstance() *mongo.Client {
	if client != nil {
		return client
	} else {
		client, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}

		if err := client.Connect(context.Background()); err != nil {
			panic(err)
		}
		if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
			panic(err)
		}
		//defer client.Disconnect(context.Background())
		fmt.Println("Database connected")
		return client
	}

}

func NewItemCollection() *mongo.Collection {
	if client != nil {
		return client.Database(database).Collection("todo-items")
	}

	client = NewDbInstance()
	return client.Database(database).Collection("todo-items")
}
