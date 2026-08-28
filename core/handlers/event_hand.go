package handlers

import (
	"event-booking/core/models"
	"event-booking/core/services"
	"event-booking/logger"
	"event-booking/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type eventHand struct {
	eventSrv services.EventService
}

func NewEventHandler(eventSrv services.EventService) *eventHand {
	return &eventHand{eventSrv: eventSrv}
}

func (h *eventHand) CreateEvent(c *fiber.Ctx) error {
	log := logger.GetLogger()
	log.Info("Create Event")

	body := models.SrvCreateEventModel{}
	if err := c.BodyParser(&body); err != nil {
		log.Error(err)
		return c.Status(400).JSON(utils.Response(400, models.InvalidRequestBodyError, nil))
	}

	res := h.eventSrv.CreateEvent(body)
	return c.Status(res.Code).JSON(res)
}

func (h *eventHand) UpdateEvent(c *fiber.Ctx) error {
	log := logger.GetLogger()
	log.Info("Update Event")

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		log.Error(err)
		return c.Status(400).JSON(utils.Response(400, models.InvalidEventIDError, nil))
	}
	body := models.SrvUpdateEventModel{}
	if err := c.BodyParser(&body); err != nil {
		log.Error(err)
		return c.Status(400).JSON(utils.Response(400, models.InvalidRequestBodyError, nil))
	}
	res := h.eventSrv.UpdateEvent(id, body)
	return c.Status(res.Code).JSON(res)
}

func (h *eventHand) DeleteEvent(c *fiber.Ctx) error {
	log := logger.GetLogger()
	log.Info("Delete Event")

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		log.Error(err)
		return c.Status(400).JSON(utils.Response(400, models.InvalidEventIDError, nil))
	}
	res := h.eventSrv.DeleteEvent(id)
	return c.Status(res.Code).JSON(res)
}
