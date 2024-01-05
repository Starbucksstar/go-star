package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"star/docs"
	"star/src/controller"
	"star/src/handler"
	"star/src/inject"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(handler.RequestCost())
	docs.SwaggerInfo.BasePath = "/api/v1"

	initUserController(router)
	initHealthController(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}

// @BasePath /health

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /health/ping [get]
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
