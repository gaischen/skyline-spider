package config

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

type Config struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	HubName   string `json:"hub_name"`
	Token     string `json:"token"`
}

var (
	ApiServerConfig Config
)

func ParserConfig() {
	filePath := "/home/api_server.conf"
	sysType := runtime.GOOS
	if sysType != "linux" { //linux环境下
		filePath = "/Users/chenhui/api_server.conf"
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Sprintf("read config file error,filePath:%v, error:%v", filePath, err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Sprintf("error decode config file to config struct, error:%v", err)
	}
	ApiServerConfig = config
	fmt.Println(config)
}
