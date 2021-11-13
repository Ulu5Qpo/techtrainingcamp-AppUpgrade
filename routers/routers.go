package routers

import (
	"AppUpgradeComb/service"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	//
	ruleConfigGroup := r.Group("ruleConfig")
	{
		ruleConfigGroup.GET("/rule", service.GetRuleList)
		ruleConfigGroup.GET("/rule/:id", service.GetRule)
		ruleConfigGroup.POST("/rule", service.AddRule)
		ruleConfigGroup.PUT("/rule/:id", service.UpdateRule)
		ruleConfigGroup.DELETE("/rule/:id", service.DeleteRule)
	}

	//session初始化
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//前端界面可以放到vies中
	r.LoadHTMLGlob("views/*")

	//登录界面
	r.GET("/login", func(c *gin.Context) {
		w := c.Writer
		t := template.Must(template.ParseFiles("views/login.html"))
		t.Execute(w, "")
	})

	//登出功能
	r.GET("/logout", service.Logout)

	//这个就是检查是否是登录状态的大概代码 如果需要判断 可以按照这个写
	r.GET("/haha", func(c *gin.Context) {
		w := c.Writer
		session := sessions.Default(c)
		username := session.Get("user")
		fmt.Println(username)
		if username == nil {
			t := template.Must(template.ParseFiles("views/login.html"))
			t.Execute(w, "")
		} else {
			c.HTML(http.StatusOK, "hahaha.html", gin.H{})
		}
	})

	//登录逻辑
	r.POST("/login", service.Login)

	r.POST("/judge", service.CheckRule)
	return r
}
