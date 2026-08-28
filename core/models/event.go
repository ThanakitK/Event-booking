package models

import "time"

type RepoCreateEventModel struct {
	Name       string    `bson:"name"`
	TotalSeats int       `bson:"total_seats"`
	StartTime  time.Time `bson:"start_time"`
}

type SrvCreateEventModel struct {
	Name       string `json:"name"`
	TotalSeats int    `json:"total_seats"`
	StartTime  string `json:"start_time"`
}

type RepoUpdateEventModel struct {
	Name       *string    `bson:"name,omitempty"`
	TotalSeats *int       `bson:"total_seats,omitempty"`
	StartTime  *time.Time `bson:"start_time,omitempty"`
	UpdatedAt  *time.Time `bson:"updated_at,omitempty"`
}

type SrvUpdateEventModel struct {
	Name       *string `json:"name"`
	TotalSeats *int    `json:"total_seats"`
	StartTime  *string `json:"start_time"`
}
