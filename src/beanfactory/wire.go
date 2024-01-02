//go:build wireinject
// +build wireinject

package beanfactory

import (
	"github.com/google/wire"
	"star/src/controller"
	"star/src/repository"
	"star/src/service"
)

var userBeanSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, controller.NewUserController)

func InitUserController() controller.UserController {
	panic(wire.Build(userBeanSet))
}

func InitUserService() service.UserService {
	panic(wire.Build(userBeanSet))
}
