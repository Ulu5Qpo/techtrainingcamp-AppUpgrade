package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/models"
	"net/http"
	"text/template"
)

func AddRule(c *gin.Context) {
	var rule models.UpgradeRule
	c.BindJSON(&rule)
	if err := models.CreateRule(&rule); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": rule,
		})
	}
}

func GetRuleList(c *gin.Context) {
	ruleList, err := models.GetAllRule()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": ruleList,
		})
	}
}

func GetRule(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	rule, err := models.GetRule(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": rule,
		})
	}
}

func UpdateRule(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	rule, err := models.GetRule(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&rule)
	if err = models.UpdateRule(rule); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": rule,
		})
	}
}

func DeleteRule(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteRule(id); err!=nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": id,
		})
	}
}

func AddUser(c *gin.Context) {
	var user models.UpgradeUser
	c.BindJSON(&user)
	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": user,
		})
	}
}

func GetUserList(c *gin.Context) {
	userList, err := models.GetAllUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": userList,
		})
	}
}

func GetUserById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": user,
		})
	}
}

func Login(c *gin.Context){
	w := c.Writer
	r := c.Request

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, _ := models.GetUser(username, password)

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

func UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&user)
	if err = models.UpdateUser(user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": user,
		})
	}
}

func DeleteUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteUser(id); err!=nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": id,
		})
	}
}