package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func RecoveryHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic recovered:", err)
				log.Println("Stack trace:", string(debug.Stack()))
				context.JSON(500, gin.H{
					"code": 500,
					"msg":  "Internal Server Error",
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
