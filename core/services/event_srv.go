package services

import (
	"event-booking/core/models"
	"event-booking/core/repositories"
	"event-booking/logger"
	"event-booking/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
		log.Error("create event failed: " + err.Error())
		return utils.Response(400, models.CreateEventError, nil)
	}

	return utils.Response(200, models.CreateEventSuccess, nil)
}

func (s *eventSrv) UpdateEvent(id primitive.ObjectID, payload models.SrvUpdateEventModel) (result models.ResponseModel) {
	log := logger.GetLogger()
	log.Info("Update Event")

	var startTime *time.Time
	if payload.StartTime != nil {
		parsed, err := time.Parse("02/01/2006", *payload.StartTime)
		if err != nil {
			log.Error("invalid start_time format: " + err.Error())
			return utils.Response(400, models.UpdateEventError, nil)
		}
		startTime = &parsed
	}

	repoPayload := models.RepoUpdateEventModel{
		Name:       payload.Name,
		TotalSeats: payload.TotalSeats,
		StartTime:  startTime,
	}

	if err := s.eventRepo.UpdateEvent(id, repoPayload); err != nil {
		log.Error("update event failed: " + err.Error())
		return utils.Response(400, models.UpdateEventError, nil)
	}

	return utils.Response(200, models.UpdateEventSuccess, nil)
}

func (s *eventSrv) DeleteEvent(id primitive.ObjectID) (result models.ResponseModel) {
	log := logger.GetLogger()
	log.Info("Delete Event")
	if err := s.eventRepo.DeleteEvent(id); err != nil {
		log.Error("delete event failed: " + err.Error())
		return utils.Response(400, models.DeleteEventError, nil)
	}
	return utils.Response(200, models.DeleteEventSuccess, nil)
}
