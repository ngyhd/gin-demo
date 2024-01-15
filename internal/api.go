package internal

import (
	"gin-demo/internal/router"
	"github.com/gin-gonic/gin"
)

func init() {

}

func Exec() *gin.Engine {
	return router.InitRouter()
}
