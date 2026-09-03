package handlers

import (
	"event-booking/core/models"
	"event-booking/core/services"
	"event-booking/logger"
	"event-booking/utils"

	"github.com/gofiber/fiber/v2"
)

type userHand struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) *userHand {
	return &userHand{userSrv: userSrv}
}

func (h *userHand) CreateUser(c *fiber.Ctx) error {
	log := logger.GetLogger()
	log.Info("Create User")

	body := models.SrvCreateUserModel{}
	if err := c.BodyParser(&body); err != nil {
		log.Error(err)
		return c.Status(400).JSON(utils.Response(400, models.InvalidRequestBodyError, nil))
	}

	res := h.userSrv.CreateUser(body)
	return c.Status(res.Code).JSON(res)
}
