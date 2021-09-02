package user_service

import (
	"github.com/phkress/mongo/models"
	userRepository "github.com/phkress/mongo/repositories/user"
)

func Create(user models.User) error {
	err := userRepository.Create(user)

	if err != nil {
		return err
	}
	return nil

}
func Read() (models.Users, error) {
	users, err := userRepository.Read()

	if err != nil {
		return nil, err
	}

	return users, nil
}
func Update(user models.User, useId string) error {
	err := userRepository.Update(user, useId)

	if err != nil {
		return err
	}
	return nil
}
func Delete(userId string) error {

	err := userRepository.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}
