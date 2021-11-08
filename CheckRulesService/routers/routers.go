package routers

import (
	"CheckRulesService/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/judge", service.CheckRule)

	return r
}
