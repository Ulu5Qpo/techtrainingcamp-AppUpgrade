package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/controller"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	ruleConfigGroup := r.Group("ruleConfig")
	{
		ruleConfigGroup.GET("/rule", controller.GetRuleList)
		ruleConfigGroup.GET("/rule/:id", controller.GetRule)
		ruleConfigGroup.POST("/rule", controller.AddRule)
		ruleConfigGroup.PUT("/rule/:id", controller.UpdateRule)
		ruleConfigGroup.DELETE("/rule/:id", controller.DeleteRule)
	}
	userConfigGroup := r.Group("userConfig")
	{
		userConfigGroup.GET("/user", controller.GetUserList)
		userConfigGroup.GET("/user/:id", controller.GetUserById)
		userConfigGroup.POST("/user", controller.AddUser)
		userConfigGroup.POST("/user/login", controller.Login)
		userConfigGroup.PUT("/user/:id", controller.UpdateUser)
		userConfigGroup.DELETE("/user/:id", controller.DeleteUser)
	}

	return r
}
