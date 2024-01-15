package logic

import (
	"gin-demo/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	// 参数校验
	// 逻辑处理
	c.JSON(http.StatusOK, response.Success(""))
}
