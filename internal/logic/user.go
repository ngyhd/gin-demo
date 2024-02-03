package logic

import (
	"gin-demo/internal/api"
	"gin-demo/internal/config"
	"gin-demo/internal/model"
	"gin-demo/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 命令行执行：
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
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode, err.Error()))
	}
	// 逻辑处理
	u := model.User{
		Username: r.UserName,
		Password: r.Password,
		Email:    r.Email,
	}
	_ = config.GetDB().Create(&u)
	c.JSON(http.StatusOK, pkg.Success())
}

func Login(c *gin.Context) {
	// 参数校验
	var r api.LoginRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode, err.Error()))
	}
	// 逻辑处理
	u := model.User{
		Username: r.UserName,
		Password: r.Password,
	}
	_ = config.GetDB().Create(&u)
	c.JSON(http.StatusOK, pkg.Success())
}

func Info(c *gin.Context) {
	// 参数校验
	var r api.InfoRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode, err.Error()))
	}
	// 逻辑处理
	u := model.User{
		Id: r.Id,
	}
	_ = config.GetDB().First(&u)
	c.JSON(http.StatusOK, pkg.Success())
}

func Update(c *gin.Context) {
	// 参数校验
	var r api.UpdateRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode, err.Error()))
	}
	// 逻辑处理
	u := model.User{
		Id:       r.Id,
		Username: r.UserName,
	}
	_ = config.GetDB().Update("username", &u)
	c.JSON(http.StatusOK, pkg.Success())
}

func Delete(c *gin.Context) {
	// 参数校验
	var r api.DeleteRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode, err.Error()))
	}
	// 逻辑处理
	u := model.User{
		Id: r.Id,
	}
	_ = config.GetDB().Delete(&u)
	c.JSON(http.StatusOK, pkg.Success())
}
