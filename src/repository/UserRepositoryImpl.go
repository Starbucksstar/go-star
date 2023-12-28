package repository

import (
	. "star/src/entity"
)

func (ur *userRepository) QueryUserByNameAndPassword(user *User) error {
	return ur.database.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error
}

func (ur *userRepository) SaveUser(user User) (uint, error) {
	if err := ur.database.Create(&user).Error; err != nil {
		return 0, err
	} else {
		return user.ID, nil
	}
}

func (ur *userRepository) DeleteUser(user User) error {
	return ur.database.Delete(&user).Error
}
