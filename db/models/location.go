package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Latitude  float64            `bson:"latitude"`
	Longitude float64            `bson:"longitude"`
	Timestamp time.Time          `bson:"timestamp"`
}
