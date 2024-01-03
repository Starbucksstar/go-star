// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"github.com/google/wire"
	"star/src/controller"
	"star/src/repository"
	"star/src/service"
)

// Injectors from wire.go:

func InitUserController() controller.UserController {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}

func InitUserService() service.UserService {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	return userService
}

// wire.go:

var UserProviderSet = wire.NewSet(repository.NewUserRepository, service.NewUserService)
