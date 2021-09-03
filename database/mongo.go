package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client       *mongo.Client
	databaseName string
}

func NewMongo(ctxMain context.Context, databaseName string) *Mongo {
	uri := "mongodb://localhost:27017/tutorial"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(ctxMain, 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return &Mongo{
		Client:       client,
		databaseName: databaseName,
	}
}
func (m *Mongo) GetCollection(collection string) *mongo.Collection {
	return m.Client.Database(m.databaseName).Collection(collection)
}
