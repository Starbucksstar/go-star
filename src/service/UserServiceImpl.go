package service

import (
	. "star/src/entity"
)

func (userService *userService) FindUser(user *User) {
	err := userService.userRepository.QueryUserByNameAndPassword(user)
	if err != nil {
		panic(err)
	}
}

func (userService *userService) SaveUser(user User) (uint, error) {
	return userService.userRepository.SaveUser(user)
}

func (userService *userService) DeleteUser(user User) error {
	return userService.userRepository.DeleteUser(user)
}
