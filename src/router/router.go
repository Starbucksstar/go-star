package router

import (
	"github.com/gin-gonic/gin"
	"star/src/controller"
	"star/src/handler"
	"star/src/inject"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(handler.RequestCost())

	initUserController(router)
	initHealthController(router)

	return router
}

func initHealthController(router *gin.Engine) {
	//Health Router
	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/ping", controller.Ping)
	}
}

func initUserController(router *gin.Engine) {
	// google/wire inject bean
	userController := inject.InitUserController()

	// User Router
	router.GET("/users", userController.Login)
	userGroup := router.Group("/users", handler.JWTAuthMiddleware())
	{
		userGroup.POST("/", userController.SignUp)
		userGroup.DELETE("/:id", userController.SignOut)
	}
}
