package repositories

import (
	"event-booking/core/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventRepository interface {
	CreateEvent(payload models.RepoCreateEventModel) (err error)

	UpdateEvent(id primitive.ObjectID, payload models.RepoUpdateEventModel) error

	DeleteEvent(id primitive.ObjectID) error
}
