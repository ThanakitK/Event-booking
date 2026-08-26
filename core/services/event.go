package services

import "event-booking/core/models"

type EventService interface {
	CreateEvent(payload models.SrvCreateEventModel) (result models.ResponseModel)
}
