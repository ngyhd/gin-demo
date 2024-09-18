package main

import (
	"gin-demo/internal"
	"gin-demo/internal/config"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// @title gin-demo API
// @version 1.0
// @description this is a sample server celler server
// @termsOfService https://www.swagger.io/terms/

// @contact.url http://www.swagger.io/support
// @contact.email abc.xyz@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9091
// @BasePath /
func main() {
	router := internal.Exec()
	s := &http.Server{
		Addr:           "0.0.0.0:" + strconv.Itoa(config.Config.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // http头部最大字节数 1024
	}
	err := s.ListenAndServe()
	if err != nil {
		zap.S().Panicf("监听失败,err:%v", err.Error())
	}
}
