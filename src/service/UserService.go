package service

import (
	. "star/src/entity"
	"star/src/repository"
)

type UserService interface {
	FindAllUser(pageNumber, pageSize int) ([]User, error)
	FindUser(user *User)
	SaveUser(user User) (uint, error)
	DeleteUser(user User) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(up repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: up,
	}
}
