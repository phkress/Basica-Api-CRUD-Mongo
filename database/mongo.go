package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongo() *Mongo {
	uri := fmt.Sprintf("mongodb://localhost:27017/tutorial")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}
	return &Mongo{
		Client: client,
	}
}
