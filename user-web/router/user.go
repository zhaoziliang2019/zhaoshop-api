package router

import (
	"github.com/gin-gonic/gin"
	"zhaoshop-api/user-web/api"
	"zhaoshop-api/user-web/middlewares"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user") //.Use(middlewares.JWTAuth())
	UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdmin(), api.GetUserList)
	UserRouter.POST("pwd_login", api.PassWordLogin)
	UserRouter.POST("register", api.RegisterUser)
	UserRouter.PATCH("detail", middlewares.JWTAuth(), api.UpdateUser)
}
