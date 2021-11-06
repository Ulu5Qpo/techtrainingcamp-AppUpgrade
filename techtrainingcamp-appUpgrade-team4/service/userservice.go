package service

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"text/template"
)
import "techtrainingcamp-appUpgrade-team4/dao"

func Login(c *gin.Context) {
	w := c.Writer
	r := c.Request

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")


	user, _ := dao.QueryForOne(username, password)

	if user.ID > 0 {
		fmt.Println("success")

		session := sessions.Default(c)
		session.Set("user", username)
		session.Save()
		fmt.Println(session.Get("user"))

	} else {
		t := template.Must(template.ParseFiles("views/login.html"))
		t.Execute(w, "用户名或密码错误")
	}
}
