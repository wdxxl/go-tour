package api

import (
	"fmt"

	"../config"
)

var (
	timeoutInMs int64
	apiEnabled  = true
)

//IsEnabled ...
func IsEnabled() bool {
	return apiEnabled
}

func init() {
	fmt.Println("init api start")
	if config.Config.ApiConf == nil {
		apiEnabled = false
		fmt.Println("init api end - return")
		return
	}

	timeoutInMs = config.Config.ApiConf.TimeoutInMs
	fmt.Println("init api end")
}
