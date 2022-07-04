package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhaoshop-api/goods-web/middlewares"
	router2 "zhaoshop-api/goods-web/router"
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
	ApiGroup := router.Group("/g/v1")
	//商品router
	router2.InitGoodsRouter(ApiGroup)
	//商品分类
	router2.InitCategoryRouter(ApiGroup)
	//品牌
	router2.InitBannerRouter(ApiGroup)
	return router
}
