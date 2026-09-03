package repositories

import (
	"context"
	"event-booking/core/models"
	"event-booking/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) UserRepository {
	return &userRepo{
		db:         db,
		collection: collection,
	}
}

func (r *userRepo) CreateUser(payload models.RepoCreateUserModel) error {
	log := logger.GetLogger()
	log.Info("Create User")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now()
	doc := bson.M{
		"name":          payload.Name,
		"user":          payload.User,
		"password_hash": payload.PasswordHash,
		"created_at":    now,
		"updated_at":    now,
	}
	_, err := r.db.Collection(r.collection).InsertOne(ctx, doc)
	return err
}
