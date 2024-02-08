package middleware

import (
	"gin-demo/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		//var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})
		if err != nil {
			zap.S().Errorf("jwt Parse err:%v", err)
			c.JSON(http.StatusOK, pkg.Fail(pkg.UserTokenErrCode))
			c.Abort()
			return
		}
		switch {
		case token.Valid:
			c.Next()
		default:
			zap.S().Errorf("jwt Parse err:%+v", err)
			c.JSON(http.StatusOK, pkg.Fail(pkg.UserTokenErrCode))
			c.Abort()
			return
		}
	}
}
