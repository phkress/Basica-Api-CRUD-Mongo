package service

import (
	"github.com/phkress/mongo/model"
	repository "github.com/phkress/mongo/repositories/user"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository,
	}
}

func (u *UserService) Create(user model.User) error {
	err := u.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

// func Create(user model.User) error {
// 	err := userRepository.Create(user)

// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }
// func Read() (model.Users, error) {
// 	users, err := userRepository.Read()

// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
// func Update(user model.User, useId string) error {
// 	err := userRepository.Update(user, useId)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func Delete(userId string) error {

// 	err := userRepository.Delete(userId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
