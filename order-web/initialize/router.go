package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhaoshop-api/order-web/middlewares"
	router2 "zhaoshop-api/order-web/router"
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
	ApiGroup := router.Group("/o/v1")
	//订单router
	router2.InitOrderRouter(ApiGroup)
	//购物车router
	router2.InitShopCartRouter(ApiGroup)
	return router
}
