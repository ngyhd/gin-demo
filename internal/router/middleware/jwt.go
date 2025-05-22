package middleware

import (
	"fmt"
	"gin-demo/pkg"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

const Secret = "AllYourBase"

type JWTInfo struct {
	privateKey []byte
	publicKey  []byte
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		//tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.nKqRzljFfJKlotnxH8auq7ui3jlIZVxI16VZQ0G0yVY"

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

func NewJWT() *JWTInfo {
	var err error
	var privateKey []byte
	var publicKey []byte

	// 读取私钥，注意路径问题
	privateKey, err = os.ReadFile("./internal/router/middleware/private.key")
	if err != nil {
		log.Fatalf("failed to load private key: %v", err)
	}
	// 读取公钥
	publicKey, err = os.ReadFile("./internal/router/middleware/public.key")
	if err != nil {
		log.Fatalf("failed to load public key: %v", err)
	}
	return &JWTInfo{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// 签发 JWT
func (j *JWTInfo) GenerateJWT(claims jwt.MapClaims) (string, error) {
	// 解析 RSA 私钥
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// 使用私钥签名
	tokenString, err := token.SignedString(rsaPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

// 解析 JWT
func (j *JWTInfo) ParseToken(tokenString string) (jwt.MapClaims, error) {
	// 解析 RSA 公钥
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// 验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确认使用的是 RS256 签名方法
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	// token 有效性校验
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Token is valid! Claims: %v\n", claims)
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

// 验证 JWT
func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			zap.S().Errorf("jwt Parse err:%v", err)
			c.JSON(http.StatusUnauthorized, pkg.Fail(pkg.UserTokenErrCode))
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
