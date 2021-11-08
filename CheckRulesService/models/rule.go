package models

type Rule struct {
	Aid                  int    `json:"aid"`
	Platform             string `json:"platform"`
	DownloadUrl          string `json:"download_url"`
	UpdateVersionCode    string `json:"update_version_code"`
	Md5                  string `json:"md5"`
	DeviceIdList         string `json:"device_id_list"`
	MaxUpdateVersionCode string `json:"max_update_version_code"`
	MinUpdateVersionCode string `json:"min_update_version_code"`
	MaxOsApi             int    `json:"max_os_api"`
	MinOsApi             int    `json:"min_os_api"`
	CpuArch              string `json:"cpu_arch"`
	Channel              string `json:"channel"`
	Title                string `json:"title"`
	UpdateTips           string `json:"update_tips"`
}

func GetAllRules() (Rules []*Rule) {
	return Rules
}
