package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhaoshop-api/user-web/middlewares"
	router2 "zhaoshop-api/user-web/router"
)

func Routers() *gin.Engine {
	router := gin.Default()
	//健康检查
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	//配置跨域
	router.Use(middlewares.Cors())
	//全局router
	ApiGroup := router.Group("u/v1")
	//用户router
	router2.InitUserRouter(ApiGroup)
	//验证码router
	router2.InitBaseRouter(ApiGroup)
	return router
}
