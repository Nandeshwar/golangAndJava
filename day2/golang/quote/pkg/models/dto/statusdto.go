package dto

import (
	"time"
)

type Status struct {
	AppName      string
	AppVer       string
	AppStartTime time.Time
}

func NewStatus(appName, appVer string, appStartTime time.Time) *Status {
	return &Status{
		AppName:      appName,
		AppVer:       appVer,
		AppStartTime: appStartTime,
	}
}
