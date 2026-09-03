package services

import "event-booking/core/models"

type UserService interface {
	CreateUser(payload models.SrvCreateUserModel) models.ResponseModel
}
