package logic

import (
	"errors"
	"gin-demo/internal/api"
	"gin-demo/internal/cache"
	"gin-demo/internal/config"
	"gin-demo/internal/model"
	"gin-demo/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	// 参数校验
	var r api.RegisterRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}
	// 逻辑处理
	u := model.User{
		Username: r.UserName,
		Password: r.Password,
		Email:    r.Email,
	}

	user := model.User{}
	//1.如果存在相同用户名则返回失败
	tx := config.GetDB().Where("username = ?", u.Username).First(&user)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		zap.S().Errorf("Register query user:%v err:%v", u.Username, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	if user.Id != 0 {
		c.JSON(http.StatusOK, pkg.Fail(pkg.UserExistsErrCode))
	}

	//2.如果存在相同的电子邮箱则返回失败
	tx = config.GetDB().Where("email = ?", u.Email).First(&user)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		zap.S().Errorf("Register queru user email:%v err:%v", u.Email, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	if user.Id != 0 {
		c.JSON(http.StatusOK, pkg.Fail(pkg.UserExistsErrCode))
	}

	tx = config.GetDB().Create(&u)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		zap.S().Errorf("Register Create user:%+v err:%v", u, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	c.JSON(http.StatusOK, pkg.Success())
}

func Login(c *gin.Context) {
	// 参数校验
	var r api.LoginRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}
	// 逻辑处理
	// 1.查询用户是否在数据库 是则登录成功，否则返回失败
	u := model.User{
		Username: r.UserName,
		Password: r.Password,
	}
	tx := config.GetDB().Where("username = ?", u.Username).First(&u)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		zap.S().Errorf("Login query user:%+v err:%v", u, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	if u.Id != 0 {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "bar",
			"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})
		mySigningKey := []byte("AllYourBase")

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err2 := token.SignedString(mySigningKey)
		if err2 != nil {
			zap.S().Errorf("Login SignedString  err:%v", err2)
			c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		}
		c.JSON(http.StatusOK, pkg.SuccessWithData(tokenString))
		return
	} else {
		c.JSON(http.StatusOK, pkg.Fail(pkg.UserNotFoundErrCode))
		return
	}
}

func Info(c *gin.Context) {
	// 参数校验
	var r api.InfoRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}
	//查询缓存
	u, err := cache.GetUserInfo(c.Request.Context(), strconv.Itoa(r.Id))
	if !errors.Is(err, redis.Nil) {
		zap.S().Errorf("Info.cache.GetUserInfo  userId:%+v err:%v", r.Id, err)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	if u != nil {
		c.JSON(http.StatusOK, pkg.SuccessWithData(u))
		return
	}

	// 逻辑处理
	u, err = cache.RefreshUserInfo(c.Request.Context(), strconv.Itoa(r.Id))
	if err != nil {
		zap.S().Errorf("Info.refreshUserInfoCache  user:%+v err:%v", r.Id, err)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
	}

	c.JSON(http.StatusOK, pkg.SuccessWithData(u))
}

func Update(c *gin.Context) {
	// 参数校验
	var r api.UpdateRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}
	// 逻辑处理
	u := model.User{
		Id:       r.Id,
		Username: r.UserName,
	}

	tx := config.GetDB().Model(&u).Update("username", u.Username)
	if tx.Error != nil {
		zap.S().Errorf("Update  user:%+v err:%v", u, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}
	// 刷新缓存
	_, err = cache.RefreshUserInfo(c.Request.Context(), strconv.Itoa(u.Id))
	if err != nil {
		zap.S().Errorf("Update.refreshUserInfoCache  user:%+v err:%v", u, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}

	c.JSON(http.StatusOK, pkg.Success())
}

func Delete(c *gin.Context) {
	// 参数校验
	var r api.DeleteRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}

	// 逻辑处理
	u := model.User{
		Id: r.Id,
	}
	tx := config.GetDB().Where("id = ?", u.Id).Delete(&u)
	if tx.Error != nil {
		zap.S().Errorf("Delete  userId:%v err:%v", u.Id, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	err = cache.DelUserInfo(c.Request.Context(), strconv.Itoa(u.Id))
	if err != nil {
		zap.S().Errorf("Delete.DelUserInfo  userId:%v err:%v", u.Id, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	c.JSON(http.StatusOK, pkg.Success())
	return
}
