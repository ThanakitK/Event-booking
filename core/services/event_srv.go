package services

import (
	"event-booking/core/models"
	"event-booking/core/repositories"
	"event-booking/logger"
	"event-booking/utils"
	"time"
)

type eventSrv struct {
	eventRepo repositories.EventRepository
}

func NewEventService(eventRepo repositories.EventRepository) EventService {
	return &eventSrv{eventRepo: eventRepo}
}

func (s *eventSrv) CreateEvent(payload models.SrvCreateEventModel) (result models.ResponseModel) {
	log := logger.GetLogger()
	log.Info("Create Event")
	startTime, _ := time.Parse("02/01/2006", payload.StartTime)
	payloadRepo := models.RepoCreateEventModel{
		Name:       payload.Name,
		TotalSeats: payload.TotalSeats,
		StartTime:  startTime,
	}
	err := s.eventRepo.CreateEvent(payloadRepo)
	if err != nil {
		return utils.Response(400, models.CreateEventError, nil)
	}

	return utils.Response(200, models.CreateEventSuccess, nil)
}
