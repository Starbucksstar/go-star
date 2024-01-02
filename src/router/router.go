package router

import (
	"github.com/gin-gonic/gin"
	"star/src/beanfactory"
	. "star/src/controller"
	"star/src/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(handler.RequestCost())

	// google/wire inject bean
	userController := beanfactory.InitUserController()

	// User Router
	router.GET("/users", userController.Login)
	userGroup := router.Group("/users", handler.JWTAuthMiddleware())
	{
		userGroup.POST("/", userController.SignUp)
		userGroup.DELETE("/:id", userController.SignOut)
	}

	//Health Router
	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/ping", Ping)
	}
	return router
}
