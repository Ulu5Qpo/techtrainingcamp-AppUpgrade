package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySql() (err error) {
	//修改为自己的数据库密码
	dsn := "root:112757995@(127.0.0.1:3306)/appUpgrade?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
