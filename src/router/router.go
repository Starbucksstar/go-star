package router

import (
	"github.com/gin-gonic/gin"
	. "star/src/controller"
	"star/src/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(handler.RequestCost())

	routerGroup := router.Group("/users", handler.JWTAuthMiddleware())
	{
		routerGroup.POST("/", SignUp)
		routerGroup.DELETE("/:id", SignOut)
	}

	router.GET("/users", Login)

	return router
}
