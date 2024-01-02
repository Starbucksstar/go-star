package repository

import (
	"gorm.io/gorm"
	. "star/src/entity"
	"star/src/global"
)

type UserRepository interface {
	FindAllUserByPage(pageNumber, pageSize int) ([]User, error)
	QueryUserByNameAndPassword(user *User) error
	SaveUser(user User) (uint, error)
	DeleteUser(user User) error
}

type UserRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		database: global.GlobalMysqlClient,
	}
}
