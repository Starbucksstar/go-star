package repository

import . "star/src/entity"
import "star/src/global"

func QueryUserByNameAndPassword(user *User) error {
	return global.GlobalMysqlClient.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error
}

func SaveUser(user User) (uint, error) {
	if err := global.GlobalMysqlClient.Create(&user).Error; err != nil {
		return 0, err
	} else {
		return user.ID, nil
	}
}

func DeleteUser(user User) error {
	return global.GlobalMysqlClient.Delete(&user).Error
}
