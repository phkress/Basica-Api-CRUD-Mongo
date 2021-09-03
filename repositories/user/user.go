package repository

import (
	"github.com/phkress/mongo/model"
)

type UserRepository interface {
	Create(user model.User) error
	//	Read() (model.Users, error)
	//	Update(user model.User, userId string) error
	//	Delete(userID string) error
}

// func Create(user model.User) error {

// 	var err error
// 	_, err = collection.InsertOne(ctx, user)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func Read() (model.Users, error) {
// 	var users model.Users

// 	filter := bson.D{}

// 	cur, err := collection.Find(ctx, filter)

// 	if err != nil {
// 		return nil, err
// 	}

// 	for cur.Next(ctx) {
// 		var user model.User

// 		err = cur.Decode(&user)

// 		if err != nil {
// 			return nil, err
// 		}

// 		users = append(users, &user)
// 	}

// 	return users, nil
// }
// func Update(user model.User, userId string) error {

// 	var err error

// 	oid, _ := primitive.ObjectIDFromHex(userId)

// 	filter := bson.M{"_id": oid}

// 	update := bson.M{
// 		"$set": bson.M{
// 			"name":       user.Name,
// 			"email":      user.Email,
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	_, err = collection.UpdateOne(ctx, filter, update)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func Delete(userId string) error {

// 	var err error
// 	var oid primitive.ObjectID

// 	oid, err = primitive.ObjectIDFromHex(userId)

// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.M{"_id": oid}

// 	_, err = collection.DeleteOne(ctx, filter)

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
