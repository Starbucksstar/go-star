package service

import (
	. "star/src/entity"
	"star/src/repository"
)

type UserService interface {
	FindUser(user *User)
	SaveUser(user User) (uint, error)
	DeleteUser(user User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(up repository.UserRepository) UserService {
	return &userService{
		userRepository: up,
	}
}
