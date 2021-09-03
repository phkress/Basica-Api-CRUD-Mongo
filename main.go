package main

import (
	"context"

	"github.com/phkress/mongo/database"
	"github.com/phkress/mongo/model"
	repository "github.com/phkress/mongo/repositories/user"
	service "github.com/phkress/mongo/services/user"
)

func main() {
	ctx := context.Background()
	mongo := database.NewMongo(ctx, "teste")
	collection := mongo.GetCollection("user")
	userRepo := repository.NewUserMongo(collection)
	userService := service.NewUserService(userRepo)

	user := model.User{}
	userService.Create(user)

}
