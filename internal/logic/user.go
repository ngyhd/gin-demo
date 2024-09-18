package logic

import (
	"errors"
	"gin-demo/internal/api"
	"gin-demo/internal/cache"
	"gin-demo/internal/config"
	"gin-demo/internal/model"
	"gin-demo/internal/router/middleware"
	"gin-demo/pkg"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// RegisterHandler
// @Summary 注册用户
// @Description Register
// @Tags 用户系统
// @Accept json
// @Produce json
// @Param data body api.RegisterRequest true "用户注册请求数据"
// @Success 200 {string} string "{"msg": "success"}"
// @Failure 400 {string} string "{"msg": "fail"}"
// @Router /register [POST]
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
		Password: pkg.HashPassword(r.Password),
		Email:    r.Email,
	}

	user := model.User{}
	//1.如果存在相同用户名则返回失败
	tx := config.DB.Where("username = ?", u.Username).First(&user)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		zap.S().Errorf("Register query user:%v err:%v", u.Username, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}

	if tx.RowsAffected != 0 { // 用户已存在
		c.JSON(http.StatusBadRequest, pkg.Fail(pkg.UserExistsErrCode))
		return
	}

	//2.如果存在相同的电子邮箱则返回失败
	tx = config.DB.Where("email = ?", u.Email).First(&user)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		zap.S().Errorf("Register queru user email:%v err:%v", u.Email, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	if tx.RowsAffected != 0 { // 邮箱已存在
		c.JSON(http.StatusOK, pkg.Fail(pkg.UserEmailExistsErrCode))
		return
	}

	tx = config.DB.Create(&u)
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
	}
	tx := config.DB.Where("username = ?", u.Username).First(&u)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		zap.S().Errorf("Login query user:%+v err:%v", u, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}

	// 用户是否存在
	if u.ID == 0 {
		c.JSON(http.StatusOK, pkg.Fail(pkg.RecordNotFoundErrCode))
		return
	}

	// 密码校验
	if pkg.CheckPassword(u.Password, r.Password) != nil {
		c.JSON(http.StatusOK, pkg.Fail(pkg.UserPasswordErrCode))
		return
	}

	j := middleware.NewJWT()
	claims := jwt.MapClaims{
		"sub":  u.ID,                                  // 用户 ID
		"name": u.Username,                            // 用户名
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // 过期时间
	}
	token, err := j.GenerateJWT(claims)
	if err != nil {
		zap.S().Infof("[CreateToken] 生成token失败")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	c.JSON(http.StatusOK, pkg.SuccessWithData(token))
}

func Info(c *gin.Context) {
	// 根据 jwt，取出当前用户信息
	claims, _ := c.Get("claims")
	currentUser := claims.(jwt.MapClaims)
	userId := currentUser["sub"].(float64)
	if userId == 0 {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}

	//查询缓存
	u, err := cache.GetUserInfo(c.Request.Context(), strconv.Itoa(int(userId)))
	if !errors.Is(err, redis.Nil) && err != nil {
		zap.S().Errorf("Info.cache.GetUserInfo  userId:%+v err:%v", userId, err)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	if u != nil {
		c.JSON(http.StatusOK, pkg.SuccessWithData(u))
		return
	}

	// 逻辑处理
	u, err = cache.RefreshUserInfo(c.Request.Context(), strconv.Itoa(int(userId)))
	if err != nil {
		zap.S().Errorf("Info.refreshUserInfoCache  user:%+v err:%v", userId, err)
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
	// 根据 jwt，取出当前用户信息
	claims, _ := c.Get("claims")
	currentUser := claims.(jwt.MapClaims)
	userId := currentUser["sub"].(float64)
	if userId == 0 {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}

	// 逻辑处理
	u := model.User{}
	config.DB.First(&u, userId)

	tx := config.DB.Model(&u).Update("username", r.UserName)
	if tx.Error != nil {
		zap.S().Errorf("Update  user:%+v err:%v", u, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}
	// 刷新缓存
	_, err = cache.RefreshUserInfo(c.Request.Context(), strconv.Itoa(int(u.ID)))
	if err != nil {
		zap.S().Errorf("Update.refreshUserInfoCache  user:%+v err:%v", u, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}

	c.JSON(http.StatusOK, pkg.SuccessWithData(u))
}

func Delete(c *gin.Context) {
	// 根据 jwt，取出当前用户信息
	claims, _ := c.Get("claims")
	currentUser := claims.(jwt.MapClaims)
	userId := currentUser["sub"].(float64)
	if userId == 0 {
		c.JSON(http.StatusOK, pkg.Fail(pkg.ParamsErrCode))
		return
	}

	// 逻辑处理
	u := model.User{}
	config.DB.First(&u, userId)
	tx := config.DB.Delete(&u)
	if tx.Error != nil {
		zap.S().Errorf("Delete  userId:%v err:%v", u.ID, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	err := cache.DelUserInfo(c.Request.Context(), strconv.Itoa(int(u.ID)))
	if err != nil {
		zap.S().Errorf("Delete.DelUserInfo  userId:%v err:%v", u.ID, tx.Error)
		c.JSON(http.StatusOK, pkg.Fail(pkg.InternalErrCode))
		return
	}
	c.JSON(http.StatusOK, pkg.Success())
}
