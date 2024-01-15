package main

import (
	"gin-demo/internal"
	"net/http"
	"time"
)

func main() {
	router := internal.Exec()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
