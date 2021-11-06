package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"techtrainingcamp-appUpgrade-team4/service"
	"techtrainingcamp-appUpgrade-team4/utils"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	//session初始化
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//获取数据库连接
	utils.InitDB()

	//前端界面可以放到vies中
	r.LoadHTMLGlob("views/*")

	//登录界面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//这个就是检查是否是登录状态的大概代码 如果需要判断 可以按照这个写
	r.GET("/haha", func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("user")
		if username == nil {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		} else {
			c.HTML(http.StatusOK, "hahaha.html", gin.H{})
		}
	})

	//登录逻辑
	r.POST("/login", service.Login)

	r.Run()

}
