package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhaoshop-api/order-web/models"
)

func IsAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		claims, _ := context.Get("claims")
		customClaims := claims.(*models.CustomClaims)
		if customClaims.AuthorityId != 2 {
			context.JSON(http.StatusForbidden, gin.H{
				"msg": "你没有访问权限",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
