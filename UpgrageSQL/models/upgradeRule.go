package models

import "github.com/oldataraxia/techtrainingcamp-AppUpgrade/UpgrageConfigService/dao"

type UpgradeRule struct {
	ID                      int    `json:"id,omitempty"`
	Aid                     int    `json:"aid,omitempty"`
	Download_url            string `json:"download_url,omitempty"`
	Update_version_code     string `json:"update_version_code,omitempty"`
	Md5                     string `json:"md_5,omitempty"`
	Device_id_list          string `json:"device_id_list,omitempty"`
	Max_update_version_code string `json:"max_update_version_code,omitempty"`
	Min_update_version_code string `json:"min_update_version_code,omitempty"`
	Max_os_api              int    `json:"max_os_api,omitempty"`
	Min_os_api              int    `json:"min_os_api,omitempty"`
	Cpu_arch                string `json:"cpu_arch,omitempty"`
	Channal                 string `json:"channal,omitempty"`
	title                   string `json:"title,omitempty"`
	update_tips             string `json:"update_tips,omitempty"`
}

func CreateRule(rule *UpgradeRule) (err error){
	err = dao.DB.Create(&rule).Error
	return err
}

func DeleteRule(id string) (err error){
	err = dao.DB.Where("id=?", id).Delete(&UpgradeRule{}).Error
	return err
}

func UpdateRule(rule *UpgradeRule) (err error) {
	err = dao.DB.Save(rule).Error
	return err
}

func GetAllRule() (ruleList []*UpgradeRule, err error) {
	if err = dao.DB.Find(&ruleList).Error; err != nil {
		return nil, err
	}
	return ruleList, nil
}

func GetRule(id string) (rule *UpgradeRule, err error) {
	rule = new(UpgradeRule)
	if err = dao.DB.Where("id=?", id).First(rule).Error; err != nil {
		return nil, err
	}
	return rule, nil
}