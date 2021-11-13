package dao

import (
	"AppUpgradeComb/models"
	"AppUpgradeComb/utils"
)

func CreateRule(rule *models.Rule) (err error) {
	err = utils.DB.Create(&rule).Error
	return err
}

func DeleteRule(id string) (err error) {
	err = utils.DB.Where("id=?", id).Delete(&models.Rule{}).Error
	return err
}

func UpdateRule(rule *models.Rule) (err error) {
	err = utils.DB.Save(rule).Error
	return err
}

func GetAllRule() (ruleList []*models.Rule, err error) {
	if err = utils.DB.Find(&ruleList).Error; err != nil {
		return nil, err
	}
	return ruleList, nil
}

func GetRuleByAid(aid int, platform string) (ruleList []*models.Rule, err error) {
	if err = utils.DB.Where("aid = ? AND platform = ?", aid, platform).Find(&ruleList).Error; err != nil {
		return nil, err
	}
	return ruleList, nil
}

func GetRule(id string) (rule *models.Rule, err error) {
	if err = utils.DB.Where("id=?", id).First(&rule).Error; err != nil {
		return nil, err
	}
	return rule, nil
}
