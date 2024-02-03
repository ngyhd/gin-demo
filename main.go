package main

import (
	"gin-demo/internal"
	"gin-demo/internal/config"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func main() {
	router := internal.Exec()
	s := &http.Server{
		Addr:           "0.0.0.0:" + strconv.Itoa(config.GetServerConfig().Port),
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
