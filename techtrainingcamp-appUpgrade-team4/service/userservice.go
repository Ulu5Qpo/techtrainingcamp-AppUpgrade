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
		session.Options(sessions.Options{MaxAge: 60 * 60})
		session.Set("user", username)
		session.Save()


	} else {
		t := template.Must(template.ParseFiles("views/login.html"))
		t.Execute(w, "用户名或密码错误")
	}
}

func Logout(c *gin.Context) {
	w := c.Writer

	session := sessions.Default(c)
	session.Clear()
	session.Save()
	t := template.Must(template.ParseFiles("views/login.html"))
	t.Execute(w, "")
}
