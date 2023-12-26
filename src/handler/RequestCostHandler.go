package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

func RequestCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Set("RequestId", uuid.New().String())
		context.Next()
		cost := time.Since(start)
		log.Println("Request=[%v]", context.Request.URL.Path, " cost=[%v]", cost)
	}
}
