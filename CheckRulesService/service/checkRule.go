package service

import (
	"CheckRulesService/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckRule(c *gin.Context) {
	var request models.ReqParameters
	c.ShouldBind(&request)
	rules := models.GetAllRules()
	for _, rule := range rules {
		hit := CheckOneRule(&request, rule)
		if hit {
			c.JSON(http.StatusOK, gin.H{
				"download_url":   rule.DownloadUrl,
				"update_version": rule.UpdateVersionCode,
				"md5":            rule.Md5,
				"title":          rule.Title,
				"update_tips":    rule.UpdateTips,
			})
			return
		}
	}

}

func CheckOneRule(request *models.ReqParameters, rule *models.Rule) bool {
	return false
}
