package dto

type ResponseStatus struct {
	AppName   string `json:"appName"`
	AppVer    string `json:"appVer"`
	AppStatus string `json:"appStatus"`
	UpTime    string `json:"upTime"`
}
