package main

import (
	"fmt"
	"quote/pkg/api"
	"quote/pkg/env"
	"time"

	"quote/pkg/models/dto"
	"quote/pkg/services/status"
)

var (
	httpPort   = env.GetIntWithDefault("HTTP_PORT", 9000)
	appName    = env.GetStringWithDefault("APP_NAME", "Quote")
	appVersion = env.GetStringWithDefault("APP_VERSION", "0.1.0")
	startTime  = time.Now()
)

func init() {
	startTime = time.Now()
}

func main() {
	fmt.Println("Welcome to world of golang1")

	statusDto := dto.NewStatus(appName, appVersion, startTime)
	statusService := status.NewService(statusDto)

	apiServer := api.NewServer(httpPort, statusService)

	apiServer.Run()

	apiServer.Close()
}
