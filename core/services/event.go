package services

import (
	"event-booking/core/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventService interface {
	CreateEvent(payload models.SrvCreateEventModel) (result models.ResponseModel)

	UpdateEvent(id primitive.ObjectID, payload models.SrvUpdateEventModel) (result models.ResponseModel)

	DeleteEvent(id primitive.ObjectID) (result models.ResponseModel)
}
