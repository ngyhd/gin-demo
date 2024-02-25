package middleware

import (
	"gin-demo/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
)

const Secret = "AllYourBase"

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.nKqRzljFfJKlotnxH8auq7ui3jlIZVxI16VZQ0G0yVY"

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(Secret), nil
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
