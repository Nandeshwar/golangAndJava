package main

import (
	"fmt"
	"quote/pkg/api"
	"quote/pkg/env"
	"time"

	"quote/pkg/models/dto"
	"quote/pkg/services/status"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Errorf("error reading config file", err)
			return
		} else {
			fmt.Errorf("error occured while reading config file", err)
			return
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	statusDto := dto.NewStatus(appName, appVersion, startTime)
	statusService := status.NewService(statusDto)

	apiServer := api.NewServer(httpPort, statusService)

	apiServer.Run()

}
