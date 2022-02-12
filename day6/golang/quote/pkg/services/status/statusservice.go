package status

import (
	"fmt"
	"time"
	
	"github.com/spf13/viper"

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
		AppStatus: viper.GetString("app.message"),
		UpTime:    fmt.Sprintf("%d %s", int64(uptime), "seconds"),
	}

}
