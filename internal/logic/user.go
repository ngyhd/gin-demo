package logic

import (
	"gin-demo/internal/api"
	"gin-demo/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// curl --location --request POST 'http://localhost:9091/register' \
// --header 'Content-Type: application/json' \
// --data-raw '{
// "username": "sean",
// "email": "8888@qq.com",
// "password": "123",
// }'
func Register(c *gin.Context) {
	// 参数校验
	var r api.RegisterRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, response.Fail(response.ParamsErrCode, err.Error()))
	}
	// 逻辑处理
	c.JSON(http.StatusOK, response.Success())
}
