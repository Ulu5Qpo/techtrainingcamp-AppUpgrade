package service

import (

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"techtrainingcamp-appUpgrade-team4/moudle"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "yellow"})
}

func Hit(c *gin.Context) {
	var respURL string

	appVersion := c.Query("appVersion")
	userDID := c.Query("userDID")

	rules := moudle.GetRules()
	for i := 0; i < len(*rules); i++ {
		if cast.ToInt(appVersion) < (*rules)[i].MinVersion || cast.ToInt(appVersion) > (*rules)[i].MaxVersion {
			continue
		}
		if cast.ToInt(userDID) < (*rules)[i].MinUserDID || cast.ToInt(userDID) > (*rules)[i].MaxUserDID {
			continue
		}

		respURL = (*rules)[i].GrayLink
	}

	c.JSON(http.StatusOK, gin.H{"respURL": respURL})
}
