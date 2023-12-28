package repository

import (
	"gorm.io/gorm"
	. "star/src/entity"
)

var _ UserRepository = (*userRepository)(nil)

type UserRepository interface {
	QueryUserByNameAndPassword(user *User) error
	SaveUser(user User) (uint, error)
	DeleteUser(user User) error
}

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		database: db,
	}
}
