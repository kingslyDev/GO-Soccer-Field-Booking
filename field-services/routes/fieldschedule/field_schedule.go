package routes

import (
	"field-services/clients"
	"field-services/constants"
	"field-services/controllers"
	"field-services/middlewares"

	"github.com/gin-gonic/gin"
)

type FieldScheduleRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IFieldScheduleRoute interface {
	Run()
}

func NewFieldScheduleRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) IFieldScheduleRoute {
	return &FieldScheduleRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (f *FieldScheduleRoute) Run() {
	group := f.group.Group("/field/schedule")
	group.GET("/lists/:uuid", middlewares.AuthenticateWithoutToken(), f.controller.GetFieldSchedule().GetAllByFieldIDAndDate)
	group.PATCH("/status", middlewares.AuthenticateWithoutToken(), f.controller.GetFieldSchedule().UpdateStatus)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, f.client),
		f.controller.GetFieldSchedule().GetAllWithPagination)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, f.client),
		f.controller.GetFieldSchedule().GetByUUID)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetFieldSchedule().Create)
	group.POST("/one-month", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetFieldSchedule().GenerateScheduleForOneMonth)
	group.PUT("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetFieldSchedule().Update)
	group.DELETE("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetFieldSchedule().Delete)
}
