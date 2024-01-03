//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"
	"star/src/controller"
	"star/src/repository"
	"star/src/service"
)

var UserProviderSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, controller.NewUserController)

func InitUserController() controller.UserController {
	panic(wire.Build(UserProviderSet))
}

func InitUserService() service.UserService {
	panic(wire.Build(UserProviderSet))
}
