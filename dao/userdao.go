package dao

import (
	"AppUpgradeComb/models"
	"AppUpgradeComb/utils"
)

func QueryForOne(username string, password string) (user *models.User, err error) {
	if err = utils.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
