package routes

import (
	"field-services/clients"
	"field-services/controllers"
	fieldRoute "field-services/routes/field"
	fieldScheduleRoute "field-services/routes/fieldschedule"
	timeRoute "field-services/routes/time"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IRegistry interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) IRegistry {
	return &Registry{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (r *Registry) fieldRoute() fieldRoute.IFieldRoute {
	return fieldRoute.NewFieldRoute(r.controller, r.group, r.client)
}

func (r *Registry) fieldScheduleRoute() fieldScheduleRoute.IFieldScheduleRoute {
	return fieldScheduleRoute.NewFieldScheduleRoute(r.controller, r.group, r.client)
}

func (r *Registry) timeRoute() timeRoute.ITimeRoute {
	return timeRoute.NewTimeRoute(r.controller, r.group, r.client)
}

func (r *Registry) Serve() {
	r.fieldRoute().Run()
	r.fieldScheduleRoute().Run()
	r.timeRoute().Run()
}
