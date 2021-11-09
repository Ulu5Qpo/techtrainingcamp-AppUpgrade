package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/config"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/dao"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/models"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/routers"
)

func main() {
	config.Init("config.ini")
	err := dao.DataBaseInit(config.Conf.MysqlConfig)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	defer dao.DBClose()
	dao.DB.AutoMigrate(&models.UpgradeRule{})
	dao.DB.AutoMigrate(&models.UpgradeUser{})

	r := routers.SetRouter()
	r.LoadHTMLGlob("views/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "login.html", " ")
	})
	if err := r.Run(fmt.Sprintf(":%d", config.Conf.Port));err!=nil {
		fmt.Printf("err:%v", err)
	}
	r.Run()
}
