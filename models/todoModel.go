package models

type Message struct {
	Data interface{} `json:"data"`
}

type Item struct {
	ItemId      string `json:"itemid,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Status      bool   `json:"status" bson:"status"`
}

type ItemsController interface {
	Create(Item) error
	Update(string, Item) error
	Delete(string) error
	GetById(string) (Item, error)
	GetAll() ([]Item, error)
}
