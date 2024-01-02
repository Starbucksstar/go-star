package repository

import (
	. "star/src/entity"
)

func (ur *UserRepositoryImpl) FindAllUserByPage(pageNumber, pageSize int) ([]User, error) {
	var users []User
	query := ur.database.Limit(pageSize).Offset((pageNumber - 1) * pageSize)
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepositoryImpl) QueryUserByNameAndPassword(user *User) error {
	return ur.database.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error
}

func (ur *UserRepositoryImpl) SaveUser(user User) (uint, error) {
	if err := ur.database.Create(&user).Error; err != nil {
		return 0, err
	} else {
		return user.ID, nil
	}
}

func (ur *UserRepositoryImpl) DeleteUser(user User) error {
	return ur.database.Delete(&user).Error
}
