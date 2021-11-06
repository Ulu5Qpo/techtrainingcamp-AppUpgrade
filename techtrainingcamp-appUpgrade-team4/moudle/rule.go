package moudle

type Rule struct {
	MinVersion int `json:"min_version"`
	MaxVersion int `json:"max_version"`
	MinUserDID int `json:"min_user_did"`
	MaxUserDID int `json:"max_user_did"`

	GrayLink string `json:"gray_link"`
}

func GetRules() *[]Rule {
	rules := []Rule{}

	rules = append(rules, Rule{
		MinVersion: 10,
		MaxVersion: 20,
		MinUserDID: 10,
		MaxUserDID: 20,
		GrayLink:   "https://www.baidu.com/",
	})

	rules = append(rules, Rule{
		MinVersion: 0,
		MaxVersion: 10,
		MinUserDID: 0,
		MaxUserDID: 10,
		GrayLink:   "https://www.baidu.com/",
	})
	return &rules
}
