package router

import (
	"github.com/gin-gonic/gin"
	"zhaoshop-api/user-web/api"
)

func InitBaseRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("base")
	{
		baseRouter.GET("captcha", api.GetCaptcha)
		baseRouter.POST("send_sms", api.SendSms)
	}
}
