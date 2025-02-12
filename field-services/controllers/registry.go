package controllers

import (
	fieldConttroller "field-services/controllers/field"
	fieldScheduleController "field-services/controllers/fieldschedule"
	timeController "field-services/controllers/time"
	"field-services/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetField() fieldConttroller.IFieldController
	GetFieldSchedule() fieldScheduleController.IFieldScheduleController
	GetTime() timeController.ITimeController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetField() fieldConttroller.IFieldController {
	return fieldConttroller.NewFieldController(r.service)
}

func (r *Registry) GetFieldSchedule() fieldScheduleController.IFieldScheduleController {
	return fieldScheduleController.NewFieldScheduleController(r.service)
}

func (r *Registry) GetTime() timeController.ITimeController {
	return timeController.NewTimeController(r.service)
}
