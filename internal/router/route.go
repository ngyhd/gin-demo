package router

import (
	"gin-demo/internal/logic"
	"gin-demo/internal/router/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "gin-demo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default()) //cors.Default() 中间件来启用CORS支持。这将允许来自任何源的GET，POST和OPTIONS请求，并允许特定的标头和方法
	router.POST("register", logic.Register)
	router.POST("login", logic.Login)
	{
		g1 := router.Group("user").Use(middleware.VerifyJWT())
		g1.POST("/info", logic.Info)
		g1.POST("/update", logic.Update)
		g1.POST("/delete", logic.Delete)
	}
	// set swagger
	// swagger.SwaggerInfo.Version = global.Conf.System.ApiVersion
	// swagger.SwaggerInfo.BasePath = v1Group.BasePath()
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.DocExpansion("none"),
		),
	)

	return router
}
