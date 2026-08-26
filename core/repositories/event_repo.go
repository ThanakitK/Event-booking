package repositories

import (
	"context"
	"event-booking/core/models"
	"event-booking/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type eventRepo struct {
	db         *mongo.Database
	collection string
}

func NewEventRepository(db *mongo.Database, collection string) EventRepository {
	return &eventRepo{
		db:         db,
		collection: collection,
	}
}

func (r *eventRepo) CreateEvent(payload models.RepoCreateEventModel) error {
	log := logger.GetLogger()
	log.Info("Create Event")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now()
	doc := bson.M{
		"name":            payload.Name,
		"total_seats":     payload.TotalSeats,
		"available_seats": payload.TotalSeats,
		"start_time":      payload.StartTime,
		"created_at":      now,
		"updated_at":      now,
	}

	_, err := r.db.Collection(r.collection).InsertOne(ctx, doc)
	return err
}
