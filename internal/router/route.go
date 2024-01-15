package router

import (
	"gin-demo/internal/logic"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("register", logic.Register)
	router.POST("login")
	router.POST("user/info")
	router.POST("user/update")
	router.POST("user/delete")
	return router
}
