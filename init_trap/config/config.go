package config

import (
	"fmt"
	"sync"
)

//Def define config struct
type Def struct {
	Host    string   `json:"host"`
	Port    int      `json:"port"`
	ApiConf *ApiConf `json:"api"`
}

//ApiConf define config struct of Api
type ApiConf struct {
	TimeoutInMs int64 `json:"timeoutInMs"`
}

var (
	Config   = &Def{}
	initOnce = &sync.Once{}
)

//Init ...
func Init() {

}

func init() {
	initOnce.Do(func() {
		fmt.Println("init do config")
		loadConfig()
	})
}

//loadConfig fake function to load config
func loadConfig() {
	Config.Host = "locahost"
	Config.Port = 8888

	Config.ApiConf = &ApiConf{
		TimeoutInMs: 1000,
	}
}
