package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/models"
	"net/http"
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