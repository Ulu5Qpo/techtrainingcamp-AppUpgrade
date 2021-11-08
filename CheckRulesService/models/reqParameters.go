package models

type ReqParameters struct {
	Version           string `json:"version"`
	DevicePlatform    string `json:"device_platform"`
	DeviceId          string `json:"device_id"`
	OsApi             int    `json:"os_api"`
	Channel           string `json:"channel"`
	VersionCode       string `json:"version_code"`
	UpdateVersionCode string `json:"update_version_code"`
	Aid               int    `json:"aid"`
	CpuArch           int    `json:"cpu_arch"`
}
