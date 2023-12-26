package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star/src/util"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusForbidden, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			context.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			context.JSON(http.StatusBadRequest, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			context.Abort()
			return
		}

		mc, err := util.ParseToken(parts[1])
		if err != nil {
			context.JSON(http.StatusBadGateway, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			context.Abort()
			return
		}

		context.Set("username", mc.Username)
		context.Set("id", mc.Id)
		context.Next()
	}
}
