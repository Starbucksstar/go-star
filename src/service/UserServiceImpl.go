package service

import (
	"log"
	. "star/src/entity"
)

func (userService *UserServiceImpl) FindAllUser(pageNumber, pageSize int) ([]User, error) {
	return userService.userRepository.FindAllUserByPage(pageNumber, pageSize)
}
func (userService *UserServiceImpl) FindUser(user *User) {
	err := userService.userRepository.QueryUserByNameAndPassword(user)
	if err != nil {
		log.Println(err)
	}
}

func (userService *UserServiceImpl) SaveUser(user User) (uint, error) {
	return userService.userRepository.SaveUser(user)
}

func (userService *UserServiceImpl) DeleteUser(user User) error {
	return userService.userRepository.DeleteUser(user)
}
