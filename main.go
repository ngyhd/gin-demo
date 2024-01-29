package main

import (
	"flag"
	"fmt"
	"gin-demo/etc"
	"gin-demo/internal"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

var ServerConfig etc.ServerConfig

func main() {
	configPathName := flag.String("config", "", "配置文件路径")
	flag.Parse()
	var configFileName string
	if *configPathName == "" {
		configFileName = "./etc/config.yaml"
	} else {
		configFileName = *configPathName
	}
	fmt.Println("配置文件路径:" + configFileName)

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("读取配置文件失败%v", err)
	}
	if err := v.Unmarshal(&ServerConfig); err != nil {
		zap.S().Panicf("解析配置文件失败 %v", err)
	}
	router := internal.Exec()
	s := &http.Server{
		Addr:           "0.0.0.0:" + strconv.Itoa(ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // http头部最大字节数 1024
	}
	s.ListenAndServe()
}
