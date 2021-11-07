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
		ruleConfigGroup.DELETE("/role/:id", controller.DeleteRule)
	}
	return r
}
