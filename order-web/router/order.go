package router

import (
	"github.com/gin-gonic/gin"
	"zhaoshop-api/order-web/api/order"
	"zhaoshop-api/order-web/api/pay"
	"zhaoshop-api/order-web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth()).Use(middlewares.Trace())
	{
		OrderRouter.GET("", order.List)                               //订单列表
		OrderRouter.POST("", order.New)                               //新建订单
		OrderRouter.GET("/:id/", middlewares.IsAdmin(), order.Detail) //订单详情
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
