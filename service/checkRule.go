package service

import (
	"AppUpgradeComb/dao"
	"AppUpgradeComb/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckRule(c *gin.Context) {
	var request models.ReqParameters
	c.ShouldBind(&request)
	rules, err := dao.GetRuleByAid(request.Aid, request.DevicePlatform)
	if err != nil {
		panic(err.Error())
	}
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
	c.JSON(http.StatusOK, gin.H{
		"msg": "do not hit",
	})

}

func CheckOneRule(request *models.ReqParameters, rule *models.Rule) bool {

	/**匹配规则
	|编号|客户端参数request|数据模型rule|匹配规则|
	|:--|---|:--|:--|
	|0唯一标识|aid|aid|request.aid == rule.aid|
	|1渠道号|channel|channel|request.channel==rule.channel|
	|2设备平台|device_platform|platform|request.device_platform==platform|
	|3cpu架构|cpu_arch|cpu_arch|request.cpu_arch == request.cpu_arch|
	|4系统版本|os_api|min_os_api，max_os_api|rule.min_os_api <= request.os_api <= rule.max_os_api|
	|5版本|update_version_code|min_update_version_code,max_update_version_code|rule.minx_update_version_code <= request.update_version_code<=rule.max_update_version_code|
	|6白名单|device_id|device_id_list|request.device_id in rule.device_id_list|
	|7未使用数据|version，version_code|None|None|
	*/

	/**操作 0, 1, 2, 3
	 * 直接比较，不同返回false
	 */

	/* 在数据库查询时已经比较过
	// 唯一标识
	if request.Aid != rule.Aid {
		return false
	}

	// 设备平台
	if request.DevicePlatform != rule.Platform {
		return false
	}
	*/

	// 1渠道号
	if request.Channel != rule.Channel {
		return false
	}

	/** 2白名单
	 * 查找是否出现在白名单 出现返回 true，否则继续比较
	 */
	for _, deviceId := range strings.Split(rule.DeviceIdList, ",") {
		if deviceId == request.DeviceId {
			return true
		}
	}

	// 3cpu架构
	if strconv.Itoa(request.CpuArch) != rule.CpuArch {
		return false
	}

	// 4系统版本
	if request.DevicePlatform != "iOS" && !(rule.MinOsApi <= request.OsApi && request.OsApi <= rule.MaxOsApi) {
		return false
	}

	// 5版本
	if !(DeviceId2IsNotGreaterDeviceId1(request.UpdateVersionCode, rule.MinUpdateVersionCode) &&
		DeviceId2IsNotGreaterDeviceId1(rule.MaxUpdateVersionCode, request.UpdateVersionCode)) {
		return false
	}

	return true
}

func DeviceId2IsNotGreaterDeviceId1(deviceId1 string, deviceId2 string) bool {
	/**
	 * input deviceId, deviceId2
	 *
	 * output : true :: deviceId >= deviceId2
	 * output : false :: deviceId = deviceId2
	 *
	 * 思路 ： 用"."分割版本号, 再补充至相同长度, 最后逐段比较
	 */
	deviceIdList1 := strings.Split(deviceId1, ".")
	deviceIdList2 := strings.Split(deviceId2, ".")

	var deviceIdLength int
	if len(deviceIdList1) > len(deviceIdList2) {
		deviceIdLength = len(deviceIdList1)
	} else {
		deviceIdLength = len(deviceIdList2)
	}
	for {
		if len(deviceIdList1) >= deviceIdLength {
			break
		}
		deviceIdList1 = append(deviceIdList1, "0")
	}
	for {
		if len(deviceIdList2) >= deviceIdLength {
			break
		}
		deviceIdList2 = append(deviceIdList2, "0")
	}

	for idx := 0; idx < deviceIdLength; idx++ {
		num1, _ := strconv.Atoi(deviceIdList1[idx])
		num2, _ := strconv.Atoi(deviceIdList2[idx])

		if num1 > num2 {
			return true
		}
		if num1 < num2 {
			return false
		}
	}

	return true
}
