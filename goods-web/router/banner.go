package router

import (
	"github.com/gin-gonic/gin"
	"zhaoshop-api/goods-web/api/banners"
	"zhaoshop-api/goods-web/middlewares"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners")
	{
		BannerRouter.GET("", banners.List)                                                        // 轮播图列表页
		BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), banners.Delete) // 删除轮播图
		BannerRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdmin(), banners.New)          //新建轮播图
		BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), banners.Update)    //修改轮播图信息
	}
}
