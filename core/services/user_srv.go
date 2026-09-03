package services

import (
	"event-booking/core/models"
	"event-booking/core/repositories"
	"event-booking/logger"
	"event-booking/utils"
)

type userSrv struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userSrv{
		userRepo: userRepo,
	}
}

func (s *userSrv) CreateUser(payload models.SrvCreateUserModel) models.ResponseModel {
	log := logger.GetLogger()
	log.Info("Create User")
	passwordHash, err := utils.HashPassword(payload.Password)
	if err != nil {
		log.Error("create user failed: " + err.Error())
		return utils.Response(400, models.CreateUserError, nil)
	}
	payloadRepo := models.RepoCreateUserModel{
		Name:         payload.Name,
		User:         payload.User,
		PasswordHash: passwordHash,
	}

	if err := s.userRepo.CreateUser(payloadRepo); err != nil {
		log.Error("create user failed: " + err.Error())
		return utils.Response(400, models.CreateUserError, nil)
	}

	return utils.Response(200, models.CreateUserSuccess, nil)
}
