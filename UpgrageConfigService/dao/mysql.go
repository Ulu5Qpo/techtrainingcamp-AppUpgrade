package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/config"
)

var (
	DB *gorm.DB
)

func DataBaseInit(cfg config.MySQLConfig) (err error) {
	mysqlMetaData := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Passwd, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", mysqlMetaData)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}

func DBClose() {
	DB.Close()
}


