package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhaoshop-api/userop-web/middlewares"
	router2 "zhaoshop-api/userop-web/router"
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
	ApiGroup := router.Group("/up/v1")
	router2.InitMessageRouter(ApiGroup)
	router2.InitUserFavRouter(ApiGroup)
	router2.InitAddressRouter(ApiGroup)
	return router
}
