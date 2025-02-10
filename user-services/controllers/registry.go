package controllers

import (
	controllers "user-services/controllers/user"
	"user-services/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetUserController() controllers.IUserController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (u *Registry) GetUserController() controllers.IUserController {
	return controllers.NewUserController(u.service)
}
