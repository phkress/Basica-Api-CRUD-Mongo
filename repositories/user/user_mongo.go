package repository

import (
	"github.com/phkress/mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	collection *mongo.Collection
}

func NewUserMongo(collection *mongo.Collection) UserRepository {
	return &UserMongo{
		collection,
	}
}

func (u *UserMongo) Create(user model.User) error {
	return nil
}
