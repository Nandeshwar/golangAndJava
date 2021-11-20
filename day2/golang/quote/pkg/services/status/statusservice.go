package status

import (
	"fmt"
	"time"

	"quote/pkg/models/dto"
)

type StatusService interface {
	GetAppStatus() *dto.ResponseStatus
}

type Status struct {
	statusDto *dto.Status
}

func NewService(statusDto *dto.Status) StatusService {
	return &Status{statusDto: statusDto}
}

func (s *Status) GetAppStatus() *dto.ResponseStatus {

	uptime := time.Since(s.statusDto.AppStartTime).Seconds()

	return &dto.ResponseStatus{
		AppName:   s.statusDto.AppName,
		AppVer:    s.statusDto.AppVer,
		AppStatus: "Up and Running",
		UpTime:    fmt.Sprintf("%d %s", int64(uptime), "seconds"),
	}

}
