package router

import (
	"github.com/gin-gonic/gin"
	"zhaoshop-api/goods-web/api/goods"
	"zhaoshop-api/goods-web/middlewares"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods").Use(middlewares.Trace())
	{
		GoodsRouter.GET("", goods.GetGoodsList)                                                    //商品列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.CreateGood)       //创建商品
		GoodsRouter.GET("/:id", goods.GetDetail)                                                   //获取商品详情
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.DeleteGood) //删除商品
		GoodsRouter.GET("/:id/stocks", goods.GetStocks)                                            //获取商品库存

		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.UpdateGoodStatus) //更新商品状态
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.UpdateGood)     //更新商品
	}
}
