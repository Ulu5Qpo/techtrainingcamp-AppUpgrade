package main

import (
	"AppUpgradeComb/models"
	"AppUpgradeComb/routers"
	"AppUpgradeComb/utils"
)

func main() {
	//
	//
	err := utils.InitMySql()
	if err != nil {
		panic(err)
	}

	utils.DB.AutoMigrate(&models.User{})
	utils.DB.AutoMigrate(&models.Rule{})

	r := routers.SetRouter()
	r.Run(":9090")
}
