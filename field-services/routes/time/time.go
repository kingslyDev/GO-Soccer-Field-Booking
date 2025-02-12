package routes

import (
	"field-services/clients"
	"field-services/constants"
	"field-services/controllers"
	"field-services/middlewares"

	"github.com/gin-gonic/gin"
)

type TimeRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type ITimeRoute interface {
	Run()
}

func NewTimeRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) ITimeRoute {
	return &TimeRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (t *TimeRoute) Run() {
	group := t.group.Group("/time")
	group.Use(middlewares.Authenticate())
	group.GET("", middlewares.CheckRole([]string{
		constants.Admin,
	}, t.client), t.controller.GetTime().GetAll)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, t.client), t.controller.GetTime().GetByUUID)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, t.client), t.controller.GetTime().Create)
}
