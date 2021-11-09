package models

import "github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/dao"

type UpgradeUser struct {
	ID                      int    `json:"id,omitempty"`
	UserName                int    `json:"username,omitempty"`
	Password                string `json:"password,omitempty"`
}

func CreateUser(user *UpgradeUser) (err error){
	err = dao.DB.Create(&user).Error
	return err
}

func DeleteUser(id string) (err error){
	err = dao.DB.Where("id=?", id).Delete(&UpgradeRule{}).Error
	return err
}

func UpdateUser(user *UpgradeUser) (err error) {
	err = dao.DB.Save(user).Error
	return err
}

func GetAllUser() (userList []*UpgradeUser, err error) {
	if err = dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func GetUser(username string, password string) (user *UpgradeUser, err error) {
	user = new(UpgradeUser)
	user.ID = 0
	if err = dao.DB.Where("username=? AND password=?", username, password).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(id string) (user *UpgradeUser, err error) {
	user = new(UpgradeUser)
	if err = dao.DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
