package repositories

import "event-booking/core/models"

type EventRepository interface {
	CreateEvent(payload models.RepoCreateEventModel) (err error)
}
