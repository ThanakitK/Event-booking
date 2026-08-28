package repositories

import (
	"context"
	"event-booking/core/models"
	"event-booking/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (r *eventRepo) UpdateEvent(id primitive.ObjectID, payload models.RepoUpdateEventModel) error {
	log := logger.GetLogger()
	log.Info("Update Event")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now()
	payload.UpdatedAt = &now

	_, err := r.db.Collection(r.collection).UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": payload},
	)
	return err
}

func (r *eventRepo) DeleteEvent(id primitive.ObjectID) error {
	log := logger.GetLogger()
	log.Info("Delete Event")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": id})
	return err
}
