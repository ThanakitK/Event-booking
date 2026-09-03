package repositories

import "event-booking/core/models"

type UserRepository interface {
	CreateUser(payload models.RepoCreateUserModel) error
}
